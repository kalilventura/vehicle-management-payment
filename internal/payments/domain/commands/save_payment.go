package commands

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

type SavePayment interface {
	Execute(payment *entities.Payment, listeners SavePaymentListeners)
}

type SavePaymentListeners struct {
	OnSuccess func(payment *entities.Payment)
	OnError   func(err error)
}
