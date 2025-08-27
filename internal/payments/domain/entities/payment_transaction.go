package entities

type PaymentTransaction struct {
	GatewayTransactionID string
	Status               string
	ResponseCode         *string
	ResponseMessage      *string
	RawResponse          []byte
}
