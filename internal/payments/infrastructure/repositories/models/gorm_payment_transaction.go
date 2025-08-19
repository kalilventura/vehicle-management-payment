package models

import (
	"time"

	"gorm.io/gorm"
)

type GormPaymentTransaction struct {
	ID                   string    `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"`
	PaymentID            string    `gorm:"type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GatewayTransactionID string    `gorm:"type:varchar(255);not null"`
	Status               string    `gorm:"type:varchar(20);not null;check:status IN ('processing','success','failure')"`
	ResponseCode         *string   `gorm:"type:varchar(50)"`
	ResponseMessage      *string   `gorm:"type:text"`
	RawResponse          []byte    `gorm:"type:jsonb"`
	CreatedAt            time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (GormPaymentTransaction) TableName() string {
	return "payment_transaction"
}

func (gv GormPaymentTransaction) BeforeCreate(_ *gorm.DB) error {
	gv.CreatedAt = time.Now().UTC()
	return nil
}
