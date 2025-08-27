package commands

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
)

type UpdatePaymentCommandStub struct {
	callback func(listeners commands.UpdatePaymentListeners)
}

func NewUpdatePaymentCommandStub() *UpdatePaymentCommandStub {
	return &UpdatePaymentCommandStub{}
}

func (stub *UpdatePaymentCommandStub) WithOnInternalServerError() *UpdatePaymentCommandStub {
	stub.callback = func(listeners commands.UpdatePaymentListeners) {
		listeners.OnError(savePaymentErr)
	}
	return stub
}

func (stub *UpdatePaymentCommandStub) WithOnSuccess() *UpdatePaymentCommandStub {
	stub.callback = func(listeners commands.UpdatePaymentListeners) {
		listeners.OnSuccess()
	}
	return stub
}

func (stub *UpdatePaymentCommandStub) Execute(_ *entities.UpdatePayment, listeners commands.UpdatePaymentListeners) {
	stub.callback(listeners)
}
