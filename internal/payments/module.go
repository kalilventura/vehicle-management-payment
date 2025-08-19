package payments

import (
	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
)

type Module struct {
	paymentControllers []entities.Controller
}

func NewModule(processPaymentController *controllers.ProcessPaymentController) *Module {
	paymentControllers := []entities.Controller{
		processPaymentController,
	}
	return &Module{paymentControllers}
}

func (m *Module) GetControllers() []entities.Controller {
	return m.paymentControllers
}
