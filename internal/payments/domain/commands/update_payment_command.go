package commands

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/repositories"
)

type UpdatePaymentCommand struct {
	repository repositories.PaymentsRepository
}

func NewUpdatePaymentCommand(repository repositories.PaymentsRepository) *UpdatePaymentCommand {
	return &UpdatePaymentCommand{
		repository,
	}
}

func (cmd *UpdatePaymentCommand) Execute(input *entities.UpdatePayment, listeners UpdatePaymentListeners) {
	if err := cmd.repository.Update(input); err != nil {
		listeners.OnError(err)
	} else {
		listeners.OnSuccess()
	}
}
