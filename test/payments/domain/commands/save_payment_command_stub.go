package commands

import (
	"errors"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
)

var savePaymentErr = errors.New("savePaymentErr")

type SavePaymentCommandStub struct {
	callback func(listeners commands.SavePaymentListeners)
}

func NewSavePaymentCommandStub() *SavePaymentCommandStub {
	return &SavePaymentCommandStub{}
}

func (stub *SavePaymentCommandStub) WithOnInternalServerError() *SavePaymentCommandStub {
	stub.callback = func(listeners commands.SavePaymentListeners) {
		listeners.OnError(savePaymentErr)
	}
	return stub
}

func (stub *SavePaymentCommandStub) WithOnSuccess(payment *entities.Payment) *SavePaymentCommandStub {
	stub.callback = func(listeners commands.SavePaymentListeners) {
		listeners.OnSuccess(payment)
	}
	return stub
}

func (stub *SavePaymentCommandStub) Execute(_ *entities.Payment, listeners commands.SavePaymentListeners) {
	stub.callback(listeners)
}
