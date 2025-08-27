package payments

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	commands.Container,
	controllers.Container,
	services.Container,
	repositories.Container,
	NewModule,
)
