package mappers

import (
	"encoding/json"
	"fmt"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/stripe/stripe-go/v82"
)

func MapToDomain(event stripe.Event) (*entities.UpdatePayment, error) {
	switch event.Type {
	case "payment_intent.succeeded":
		return buildUpdatePayment(event, "approved")
	case "payment_intent.canceled":
		return buildUpdatePayment(event, "failed")
	default:
		return nil, fmt.Errorf("payment status %s is not supported", event.Type)
	}
}

func buildUpdatePayment(event stripe.Event, status string) (*entities.UpdatePayment, error) {
	intent, err := mapPaymentIntent(event)
	if err != nil {
		return nil, err
	}
	transaction := &entities.PaymentTransaction{
		GatewayTransactionID: intent.ID,
		Status:               string(intent.Status),
		ResponseCode:         nil,
		ResponseMessage:      nil,
		RawResponse:          event.Data.Raw,
	}
	payment := &entities.UpdatePayment{
		GatewayPaymentID:   intent.ID,
		Status:             status,
		PaymentTransaction: transaction,
	}
	return payment, nil
}

func mapPaymentIntent(event stripe.Event) (*stripe.PaymentIntent, error) {
	paymentIntent := &stripe.PaymentIntent{}
	err := json.Unmarshal(event.Data.Raw, paymentIntent)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment intent. Reason: %w", err)
	}
	return paymentIntent, nil
}
