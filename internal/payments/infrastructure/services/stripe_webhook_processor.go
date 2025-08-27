package services

import (
	"fmt"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers/mappers"
	global "github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	logger "github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v82/webhook"
)

// StripeWebhookProcessor is the concrete implementation for handling Stripe webhooks.
type StripeWebhookProcessor struct {
	settings *global.WebhookSettings
}

// NewStripeWebhookProcessor creates a new instance of the Stripe webhook processor.
func NewStripeWebhookProcessor(settings *global.WebhookSettings) *StripeWebhookProcessor {
	return &StripeWebhookProcessor{
		settings: settings,
	}
}

// Process implements the WebhookProcessor interface for Stripe.
func (p *StripeWebhookProcessor) Process(payload []byte, signature string) (*entities.UpdatePayment, error) {
	event, err := webhook.ConstructEvent(payload, signature, p.settings.Secret)
	if err != nil {
		return nil, fmt.Errorf("error constructing stripe event: %w", err)
	}
	logger.Infof("Request %s received with the status: %s", event.Data.Object["id"], event.Type)

	entity, err := mappers.MapToDomain(event)
	if err != nil {
		return nil, fmt.Errorf("error mapping stripe event to domain entity: %w", err)
	}
	return entity, nil
}
