//go:build unit

package services_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/services"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/stretchr/testify/assert"
)

func TestStripePaymentService(t *testing.T) {
	t.Run("should return an error to create the payment", func(t *testing.T) {
		// given
		payment := builders.NewPaymentBuilder().Build()

		settings := &entities.PaymentSettings{
			StripeKey: gofakeit.UUID(),
		}
		service := services.NewStripePaymentService(settings)

		// then
		_, err := service.Pay(&payment)

		// then
		assert.Error(t, err)
	})
}
