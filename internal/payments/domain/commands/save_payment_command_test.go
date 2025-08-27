//go:build unit

package commands_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/kalilventura/vehicle-management-payment/test/payments/infrastructure/repositories"
	"github.com/kalilventura/vehicle-management-payment/test/payments/infrastructure/services"
	"github.com/stretchr/testify/assert"
)

func TestSavePaymentCommand(t *testing.T) {
	t.Run("should call OnSuccess when the payment was processed", func(t *testing.T) {
		// given
		input := builders.NewPaymentBuilder().Build()
		transaction := builders.NewPaymentTransactionBuilder().Build()

		repo := repositories.NewInMemoryPaymentsRepository()
		paymentGateway := services.NewPaymentServiceStub().WithOnSuccess(&transaction)
		command := commands.NewSavePaymentCommand(repo, paymentGateway)

		listeners := commands.SavePaymentListeners{
			OnSuccess: func(payment *entities.Payment) {
				assert.NotNil(t, payment)
			},
		}

		// when
		command.Execute(&input, listeners)
	})

	t.Run("should call OnError when the payment gateway was failed", func(t *testing.T) {
		// given
		input := builders.NewPaymentBuilder().Build()
		repo := repositories.NewInMemoryPaymentsRepository()
		paymentGateway := services.NewPaymentServiceStub().WithError()
		command := commands.NewSavePaymentCommand(repo, paymentGateway)

		listeners := commands.SavePaymentListeners{
			OnError: func(err error) {
				assert.Error(t, err)
			},
		}

		// when
		command.Execute(&input, listeners)
	})

	t.Run("should call OnError when there's an unexpected error on the database", func(t *testing.T) {
		// given
		input := builders.NewPaymentBuilder().Build()
		repo := repositories.NewInMemoryPaymentsRepository().WithOnError()
		paymentGateway := services.NewPaymentServiceStub()
		command := commands.NewSavePaymentCommand(repo, paymentGateway)

		listeners := commands.SavePaymentListeners{
			OnError: func(err error) {
				assert.Error(t, err)
			},
		}

		// when
		command.Execute(&input, listeners)
	})
}
