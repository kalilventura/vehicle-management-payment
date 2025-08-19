package mappers

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/models"
)

func MapToEntity(payment *entities.Payment) *models.GormPayment {
	return &models.GormPayment{
		VehicleID: payment.VehicleID,
		Cpf:       payment.Cpf,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt,
	}
}
