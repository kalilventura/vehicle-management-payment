package repositories

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

type PaymentsRepository interface {
	Save(payment *entities.Payment) error
	Update(payment *entities.UpdatePayment) error
}
