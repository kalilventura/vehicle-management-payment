package controllers

import (
	"io"
	"net/http"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/services"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/controllers"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type WebhookPaymentController struct {
	command   commands.UpdatePayment
	processor services.WebhookProcessor
}

func NewWebhookPaymentController(
	command commands.UpdatePayment,
	processor services.WebhookProcessor,
) *WebhookPaymentController {
	return &WebhookPaymentController{
		command:   command,
		processor: processor,
	}
}

func (ctrl *WebhookPaymentController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/payments/webhook",
	}
}

// Execute Update payments webhook
//
// @Summary Update payments webhook
// @Description Update payments webhook
// @BasePath /v1/payments/webhook
// @Tags payments
// @Accept application/json
// @Produce application/json
// @Param request body stripe.Event true "Request body"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /v1/payments/webhook [post]
func (ctrl *WebhookPaymentController) Execute(ectx echo.Context) error {
	body, err := io.ReadAll(ectx.Request().Body)
	if err != nil {
		logger.Errorf("An unexpected error occurred while reading the request body. Reason: %v\n", err)
		return ctrl.onInternalServerError(ectx, err)
	}
	signature := ectx.Request().Header.Get("Stripe-Signature")

	entity, err := ctrl.processor.Process(body, signature)
	if err != nil {
		return ctrl.onBadRequest(ectx, err)
	}

	var handle error
	commandListeners := commands.UpdatePaymentListeners{
		OnSuccess: func() {
			handle = ctrl.onSuccess(ectx, entity.GatewayPaymentID)
		},
		OnError: func(err error) {
			handle = ctrl.onInternalServerError(ectx, err)
		},
	}
	ctrl.command.Execute(entity, commandListeners)
	return handle
}

func (ctrl *WebhookPaymentController) onBadRequest(ectx echo.Context, err error) error {
	logger.Errorf("⚠️ Webhook processing failed. Reason: %v\n", err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		[]string{err.Error()},
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *WebhookPaymentController) onInternalServerError(ectx echo.Context, err error) error {
	logger.Errorf("🔥 An internal server error occurred. Reason: %v\n", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}

func (ctrl *WebhookPaymentController) onSuccess(ectx echo.Context, paymentID string) error {
	logger.Infof("Successfully processed webhook for payment: %s.", paymentID)
	return ectx.NoContent(http.StatusOK)
}
