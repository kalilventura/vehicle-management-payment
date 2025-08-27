//go:build !test && wireinject

package main

import (
	"os"
	"strconv"

	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management-payment/internal/payments"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/configuration"
)

func InjectModules() []entities.HTTPModule {
	wire.Build(
		injectWebhookSettings,
		injectPaymentSettings,
		injectDatabaseSettings,
		configuration.NewDatabaseClient,
		payments.Container,
		newModules,
	)
	return nil
}

func InjectSettings() *entities.Settings {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &entities.Settings{port}
}

func injectDatabaseSettings() *entities.DatabaseSettings {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbSSL := os.Getenv("DB_SSL")
	return entities.NewDatabaseSettings(
		host,
		name,
		port,
		user,
		password,
		dbSSL,
	)
}

func injectPaymentSettings() *entities.PaymentSettings {
	stripeKey := os.Getenv("STRIPE_KEY")
	return &entities.PaymentSettings{stripeKey}
}

func injectWebhookSettings() *entities.WebhookSettings {
	stripeKey := os.Getenv("STRIPE_WEBHOOK_KEY")
	return &entities.WebhookSettings{stripeKey}
}

func newModules(paymentsModule *payments.Module) []entities.HTTPModule {
	return []entities.HTTPModule{
		paymentsModule,
	}
}
