package services

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

type PaymentService interface {
	Pay(input *entities.Payment) (*entities.PaymentTransaction, error)
}
