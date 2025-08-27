package controllers

import (
	"net/http"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	entities2 "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers/requests"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers/responses"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/controllers"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type ProcessPaymentController struct {
	command commands.SavePayment
}

func NewProcessPaymentController(command commands.SavePayment) *ProcessPaymentController {
	return &ProcessPaymentController{command: command}
}

func (ctrl *ProcessPaymentController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/payments",
	}
}

// Execute Save a new payment
//
// @Summary Save a new payment
// @Description Save a new payment
// @BasePath /v1/payments
// @Tags payments
// @Accept application/json
// @Produce application/json
// @Param request body requests.ProcessPaymentRequest true "Request body"
// @Success 200 {object} controllers.SuccessResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @Router /v1/payments [post]
func (ctrl *ProcessPaymentController) Execute(ectx echo.Context) error {
	paymentRequest := new(requests.ProcessPaymentRequest)
	if err := ectx.Bind(paymentRequest); err != nil {
		return ctrl.onInvalid(ectx, err)
	}
	entity := paymentRequest.ToDomain()

	var handle error
	listeners := commands.SavePaymentListeners{
		OnSuccess: func(payment *entities2.Payment) {
			handle = ctrl.onSuccess(ectx, payment)
		},
		OnError: func(err error) {
			handle = ctrl.onError(ectx, err)
		},
	}
	ctrl.command.Execute(entity, listeners)
	return handle
}

func (ctrl *ProcessPaymentController) onInvalid(ectx echo.Context, err error) error {
	validationErrors := ctrl.extractValidationErrors(err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		validationErrors,
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *ProcessPaymentController) onSuccess(ectx echo.Context, payment *entities2.Payment) error {
	responseBody := responses.CreateNewProcessPaymentResponse(payment)
	response := controllers.NewSuccessResponse(http.StatusAccepted, responseBody)
	return ectx.JSON(http.StatusAccepted, response)
}

func (ctrl *ProcessPaymentController) onError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}

func (ctrl *ProcessPaymentController) extractValidationErrors(err error) map[string]string {
	return map[string]string{
		"generic": err.Error(),
	}
}
