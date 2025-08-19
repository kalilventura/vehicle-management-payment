package commands

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/repositories"
)

type SavePaymentCommand struct {
	repository repositories.PaymentsRepository
}

func NewSavePaymentCommand(repository repositories.PaymentsRepository) *SavePaymentCommand {
	return &SavePaymentCommand{repository}
}

func (cmd *SavePaymentCommand) Execute(payment *entities.Payment, listeners SavePaymentListeners) {
	if err := cmd.repository.Save(payment); err != nil {
		listeners.OnError(err)
	} else {
		listeners.OnSuccess(payment)
	}
}
