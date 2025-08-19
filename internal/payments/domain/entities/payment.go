package entities

import "time"

type Payment struct {
	ID        string
	Status    string
	VehicleID string
	Cpf       string
	Amount    float64
	CreatedAt time.Time
}
