package repositories

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management-payment/internal/payments/domain/repositories"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewGormPaymentsRepository,
	wire.Bind(new(repositories.PaymentsRepository), new(*GormPaymentsRepository)),
)
