package requests

import (
	"time"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
)

// ProcessPaymentRequest
// @Description Object that represents a vehicle payment
type ProcessPaymentRequest struct {
	VehicleID string    `json:"vehicle_id" binding:"required"`
	CPF       string    `json:"cpf" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
} // @name ProcessPaymentRequest

func (dto ProcessPaymentRequest) ToDomain() *entities.Payment {
	return &entities.Payment{
		VehicleID: dto.VehicleID,
		Cpf:       dto.CPF,
		Amount:    dto.Amount,
		CreatedAt: dto.CreatedAt,
	}
}
