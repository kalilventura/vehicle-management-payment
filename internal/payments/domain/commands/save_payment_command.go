package commands

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/repositories"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/services"
)

type SavePaymentCommand struct {
	repository repositories.PaymentsRepository
	gateway    services.PaymentService
}

func NewSavePaymentCommand(
	repository repositories.PaymentsRepository,
	gateway services.PaymentService) *SavePaymentCommand {
	return &SavePaymentCommand{
		repository,
		gateway,
	}
}

func (cmd *SavePaymentCommand) Execute(payment *entities.Payment, listeners SavePaymentListeners) {
	paymentTransaction, err := cmd.gateway.Pay(payment)
	if err != nil {
		listeners.OnError(err)
		return
	}

	payment.PaymentTransaction = paymentTransaction
	if saveErr := cmd.repository.Save(payment); saveErr != nil {
		listeners.OnError(saveErr)
	} else {
		listeners.OnSuccess(payment)
	}
}
