package services

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewStripePaymentService,
	wire.Bind(new(services.PaymentService), new(*StripePaymentService)),
	NewStripeWebhookProcessor,
	wire.Bind(new(services.WebhookProcessor), new(*StripeWebhookProcessor)),
)
