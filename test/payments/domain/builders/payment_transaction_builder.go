package builders

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	builders "github.com/kalilventura/vehicle-management-payment/test/shared"
)

type PaymentTransactionBuilder struct {
	builders.BaseBuilder[entities.PaymentTransaction]
}

func NewPaymentTransactionBuilder() *PaymentTransactionBuilder {
	return &PaymentTransactionBuilder{}
}

func (b *PaymentTransactionBuilder) WithRawResponse(value []byte) *PaymentTransactionBuilder {
	b.AppendModifier(func(r *entities.PaymentTransaction) {
		r.RawResponse = value
	})
	return b
}
