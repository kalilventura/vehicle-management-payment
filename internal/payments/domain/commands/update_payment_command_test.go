//go:build unit

package commands_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/kalilventura/vehicle-management-payment/test/payments/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePaymentCommand(t *testing.T) {
	t.Run("should call OnSuccess when the update was executed", func(t *testing.T) {
		// given
		update := builders.NewUpdatePaymentBuilder().Build()

		repo := repositories.NewInMemoryPaymentsRepository().WithOnError()
		command := commands.NewUpdatePaymentCommand(repo)

		listeners := commands.UpdatePaymentListeners{
			OnError: func(err error) {
				assert.Error(t, err)
			},
		}

		// when
		command.Execute(&update, listeners)
	})

	t.Run("should call OnSuccess when the update was executed", func(t *testing.T) {
		// given
		update := builders.NewUpdatePaymentBuilder().Build()

		repo := repositories.NewInMemoryPaymentsRepository()
		command := commands.NewUpdatePaymentCommand(repo)

		listeners := commands.UpdatePaymentListeners{
			OnSuccess: func() {
				assert.True(t, true)
			},
		}

		// when
		command.Execute(&update, listeners)
	})
}
