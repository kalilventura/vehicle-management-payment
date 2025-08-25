package services

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

// WebhookProcessor defines the contract for processing and validating a webhook payload.
type WebhookProcessor interface {
	Process(payload []byte, signature string) (*entities.UpdatePayment, error)
}
