package models

import (
	"time"

	"gorm.io/gorm"
)

type GormPayment struct {
	ID           string                   `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	VehicleID    string                   `gorm:"not null"`
	Cpf          string                   `gorm:"type:text;not null"`
	Amount       float64                  `gorm:"not null;check:amount > 0"`
	Status       string                   `gorm:"type:varchar(20);not null;default:'pending'"`
	CreatedAt    time.Time                `gorm:"not null;default:CURRENT_TIMESTAMP"`
	ModifiedAt   *time.Time               `gorm:"column:updated_at"`
	Transactions []GormPaymentTransaction `gorm:"foreignKey:PaymentID"`
}

func (GormPayment) TableName() string {
	return "vehicle_payment"
}

func (g GormPayment) BeforeUpdate(_ *gorm.DB) error {
	now := time.Now().UTC()
	g.ModifiedAt = &now
	return nil
}
