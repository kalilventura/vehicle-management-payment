package builders

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	builders "github.com/kalilventura/vehicle-management-payment/test/shared"
)

type PaymentBuilder struct {
	builders.BaseBuilder[entities.Payment]
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{}
}

func (b *PaymentBuilder) WithAmount(value float64) *PaymentBuilder {
	b.AppendModifier(func(r *entities.Payment) {
		r.Amount = value
	})
	return b
}

func (b *PaymentBuilder) WithPaymentTransaction(value *entities.PaymentTransaction) *PaymentBuilder {
	b.AppendModifier(func(r *entities.Payment) {
		r.PaymentTransaction = value
	})
	return b
}
