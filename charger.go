package uculqi

import "time"

type ChargeResponse struct {
	ID          string
	At          time.Time
	How         float64
	UserMessage string
}

type Charger interface {
	ExecuteCharge(invoice *Invoice) (*ChargeResponse, error)
}
