package repositories

import (
	"fmt"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/mappers"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/models"
	"gorm.io/gorm"
)

type GormPaymentsRepository struct {
	client *gorm.DB
}

func NewGormPaymentsRepository(client *gorm.DB) *GormPaymentsRepository {
	return &GormPaymentsRepository{client}
}

func (r *GormPaymentsRepository) Save(payment *entities.Payment) error {
	dbEntity := mappers.MapToEntity(payment)
	if err := r.client.Create(&dbEntity).Error; err != nil {
		return fmt.Errorf("failed to save payment. Reason: %w", err)
	}
	payment.ID = dbEntity.ID
	payment.Status = dbEntity.Status
	return nil
}

func (r *GormPaymentsRepository) Update(input *entities.UpdatePayment) error {
	dbEntity := mappers.MapUpdatePaymentToGorm(input)
	return r.client.Transaction(func(tx *gorm.DB) error {
		var payment models.GormPayment
		transactionID := dbEntity.Transactions[0].GatewayTransactionID
		err := tx.Model(&models.GormPayment{}).
			Joins("JOIN payment_transaction t ON t.payment_id = vehicle_payment.id").
			Where("t.gateway_transaction_id = ?", transactionID).
			First(&payment).Error
		if err != nil {
			return err
		}

		if updateErr := tx.
			Model(&payment).
			Update("status", dbEntity.Status).
			Error; updateErr != nil {
			return err
		}

		newTx := dbEntity.Transactions[0]
		newTx.PaymentID = payment.ID
		if createErr := tx.Create(&newTx).Error; createErr != nil {
			return createErr
		}
		return nil
	})
}
