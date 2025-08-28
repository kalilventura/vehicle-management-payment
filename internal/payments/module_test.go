//go:build unit

package payments_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	t.Run("should return the quantity of controllers successfully", func(t *testing.T) {
		// given
		paymentsController := controllers.NewProcessPaymentController(nil)
		webhookController := controllers.NewWebhookPaymentController(nil, nil)

		app := payments.NewModule(paymentsController, webhookController)

		// when
		paymentControllers := app.GetControllers()

		// then
		assert.Len(t, paymentControllers, 2)
	})
}
