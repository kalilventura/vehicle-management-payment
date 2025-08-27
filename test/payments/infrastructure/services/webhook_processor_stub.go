package services

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

type WebhookProcessorStub struct {
	err    error
	entity *entities.UpdatePayment
}

func NewWebhookProcessorStub() *WebhookProcessorStub {
	return &WebhookProcessorStub{}
}

func (s *WebhookProcessorStub) WithSuccess(entity *entities.UpdatePayment) *WebhookProcessorStub {
	s.entity = entity
	return s
}

func (s *WebhookProcessorStub) WithError() *WebhookProcessorStub {
	s.err = payServiceErr
	return s
}

func (s *WebhookProcessorStub) Process(_ []byte, _ string) (*entities.UpdatePayment, error) {
	return s.entity, s.err
}
