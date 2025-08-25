//go:build unit

package services_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/services"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestStripeWebhookProcessor(t *testing.T) {
	t.Run("should return an error to decode a webhook request", func(t *testing.T) {
		// given
		settings := &entities.WebhookSettings{
			Secret: gofakeit.UUID(),
		}
		service := services.NewStripeWebhookProcessor(settings)

		// when
		_, err := service.Process(nil, "")

		// then
		assert.Error(t, err)
	})
}
