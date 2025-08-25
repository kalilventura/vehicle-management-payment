package builders

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"
	builders "github.com/kalilventura/vehicle-management-payment/test/shared"
)

type UpdatePaymentBuilder struct {
	builders.BaseBuilder[entities.UpdatePayment]
}

func NewUpdatePaymentBuilder() *UpdatePaymentBuilder {
	return &UpdatePaymentBuilder{}
}

func (b *UpdatePaymentBuilder) WithGatewayPaymentID(value string) *UpdatePaymentBuilder {
	b.AppendModifier(func(r *entities.UpdatePayment) {
		r.GatewayPaymentID = value
	})
	return b
}

func (b *UpdatePaymentBuilder) WithPaymentTransaction(value *entities.PaymentTransaction) *UpdatePaymentBuilder {
	b.AppendModifier(func(r *entities.UpdatePayment) {
		r.PaymentTransaction = value
	})
	return b
}
