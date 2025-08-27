package repositories

import (
	"errors"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
)

var payRepositoryErr = errors.New("payRepositoryErr")

type InMemoryPaymentsRepository struct {
	err error
}

func NewInMemoryPaymentsRepository() *InMemoryPaymentsRepository {
	return &InMemoryPaymentsRepository{}
}

func (r *InMemoryPaymentsRepository) WithOnError() *InMemoryPaymentsRepository {
	r.err = payRepositoryErr
	return r
}

func (r *InMemoryPaymentsRepository) Save(_ *entities.Payment) error {
	return r.err
}

func (r *InMemoryPaymentsRepository) Update(_ *entities.UpdatePayment) error {
	return r.err
}
