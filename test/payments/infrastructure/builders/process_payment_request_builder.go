package builders

import (
	"bytes"
	"encoding/json"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers/requests"
	builders "github.com/kalilventura/vehicle-management-payment/test/shared"
)

type ProcessPaymentRequestBuilder struct {
	builders.BaseBuilder[requests.ProcessPaymentRequest]
}

func NewProcessPaymentRequestBuilder() *ProcessPaymentRequestBuilder {
	return &ProcessPaymentRequestBuilder{}
}

func (b *ProcessPaymentRequestBuilder) WithValidData() *ProcessPaymentRequestBuilder {
	return b.WithCPF("12321321312312").WithVehicleID(gofakeit.UUID()).WithAmount(gofakeit.Float64())
}

func (b *ProcessPaymentRequestBuilder) WithCPF(value string) *ProcessPaymentRequestBuilder {
	b.AppendModifier(func(r *requests.ProcessPaymentRequest) {
		r.CPF = value
	})
	return b
}

func (b *ProcessPaymentRequestBuilder) WithAmount(value float64) *ProcessPaymentRequestBuilder {
	b.AppendModifier(func(r *requests.ProcessPaymentRequest) {
		r.Amount = value
	})
	return b
}

func (b *ProcessPaymentRequestBuilder) WithVehicleID(value string) *ProcessPaymentRequestBuilder {
	b.AppendModifier(func(r *requests.ProcessPaymentRequest) {
		r.VehicleID = value
	})
	return b
}

func (b *ProcessPaymentRequestBuilder) BuildRequest() *bytes.Buffer {
	data := b.Build()
	requestBodyBytes, _ := json.Marshal(data)
	return bytes.NewBuffer(requestBodyBytes)
}
