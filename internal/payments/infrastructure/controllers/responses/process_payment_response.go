package responses

import "github.com/kalilventura/vehicle-management-payment/internal/payments/domain/entities"

// ProcessPaymentResponse
// @Description Object that represents a vehicle payment response
type ProcessPaymentResponse struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
} // @name ProcessPaymentResponse

func CreateNewProcessPaymentResponse(payment *entities.Payment) ProcessPaymentResponse {
	return ProcessPaymentResponse{
		PaymentID: payment.ID,
		Status:    payment.Status,
	}
}
