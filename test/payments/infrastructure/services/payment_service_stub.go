package services

import (
	"errors"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
)

var payServiceErr = errors.New("payServiceErr")

type PaymentServiceStub struct {
	paymentTransaction *entities.PaymentTransaction
	err                error
}

func NewPaymentServiceStub() *PaymentServiceStub {
	return &PaymentServiceStub{}
}

func (s *PaymentServiceStub) WithError() *PaymentServiceStub {
	s.err = payServiceErr
	return s
}

func (s *PaymentServiceStub) WithOnSuccess(entity *entities.PaymentTransaction) *PaymentServiceStub {
	s.paymentTransaction = entity
	return s
}

func (s *PaymentServiceStub) Pay(_ *entities.Payment) (*entities.PaymentTransaction, error) {
	return s.paymentTransaction, s.err
}
