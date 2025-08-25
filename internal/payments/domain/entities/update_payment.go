package entities

type UpdatePayment struct {
	GatewayPaymentID   string
	Status             string
	PaymentTransaction *PaymentTransaction
}
