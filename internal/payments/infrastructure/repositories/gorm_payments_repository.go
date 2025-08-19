package repositories

import (
	"fmt"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/mappers"
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
