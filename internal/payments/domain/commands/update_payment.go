package commands

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

type UpdatePayment interface {
	Execute(input *entities.UpdatePayment, listeners UpdatePaymentListeners)
}

type UpdatePaymentListeners struct {
	OnSuccess func()
	OnError   func(err error)
}
