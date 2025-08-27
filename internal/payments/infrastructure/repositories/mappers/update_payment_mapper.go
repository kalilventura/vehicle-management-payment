package mappers

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/models"
)

func MapUpdatePaymentToGorm(input *entities.UpdatePayment) *models.GormPayment {
	transaction := models.GormPaymentTransaction{
		GatewayTransactionID: input.PaymentTransaction.GatewayTransactionID,
		Status:               input.PaymentTransaction.Status,
		ResponseCode:         input.PaymentTransaction.ResponseCode,
		ResponseMessage:      input.PaymentTransaction.ResponseMessage,
		RawResponse:          input.PaymentTransaction.RawResponse,
	}
	payment := &models.GormPayment{
		Status:       input.Status,
		Transactions: []models.GormPaymentTransaction{transaction},
	}
	return payment
}
