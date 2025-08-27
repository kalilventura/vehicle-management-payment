package entities

type Payment struct {
	ID                 string
	Status             string
	VehicleID          string
	Cpf                string
	PaymentTransaction *PaymentTransaction
	Amount             float64
}
