package services

import (
	"fmt"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	global "github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	logger "github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

type StripePaymentService struct {
	key string
}

func NewStripePaymentService(settings *global.PaymentSettings) *StripePaymentService {
	return &StripePaymentService{
		settings.StripeKey,
	}
}

func (s *StripePaymentService) Pay(input *entities.Payment) (*entities.PaymentTransaction, error) {
	stripe.Key = s.key
	amount := int64(input.Amount)

	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String("brl"),
		Description: stripe.String("Payment for vehicle"),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled:        stripe.Bool(true),
			AllowRedirects: stripe.String("never"),
		},
	}

	result, err := paymentintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent. Reason: %w", err)
	}
	respStatus := string(result.Status)
	logger.Infof("Payment: %v with the Status: %s",
		result.ID,
		respStatus)

	transaction := &entities.PaymentTransaction{
		GatewayTransactionID: result.ID,
		Status:               respStatus,
		ResponseCode:         &result.APIResource.LastResponse.Status,
		ResponseMessage:      &respStatus,
		RawResponse:          result.APIResource.LastResponse.RawJSON,
	}
	return transaction, nil
}
