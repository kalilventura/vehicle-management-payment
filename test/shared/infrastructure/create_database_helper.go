package infrastructure

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories/models"
	"gorm.io/gorm"
)

func CreateDatabaseStructure(db *gorm.DB) error {
	err := db.AutoMigrate(models.GormPayment{})
	err = db.AutoMigrate(models.GormPaymentTransaction{})
	return err
}
