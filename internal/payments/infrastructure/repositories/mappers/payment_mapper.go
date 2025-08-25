package mappers

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/models"
)

func MapToEntity(payment *entities.Payment) *models.GormPayment {
	transaction := models.GormPaymentTransaction{
		GatewayTransactionID: payment.PaymentTransaction.GatewayTransactionID,
		Status:               payment.PaymentTransaction.Status,
		ResponseCode:         payment.PaymentTransaction.ResponseCode,
		ResponseMessage:      payment.PaymentTransaction.ResponseMessage,
		RawResponse:          payment.PaymentTransaction.RawResponse,
	}
	return &models.GormPayment{
		VehicleID:    payment.VehicleID,
		Cpf:          payment.Cpf,
		Amount:       payment.Amount,
		Transactions: []models.GormPaymentTransaction{transaction},
	}
}
