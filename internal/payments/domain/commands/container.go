package commands

import (
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSavePaymentCommand,
	wire.Bind(new(SavePayment), new(*SavePaymentCommand)),
	NewUpdatePaymentCommand,
	wire.Bind(new(UpdatePayment), new(*UpdatePaymentCommand)),
)
