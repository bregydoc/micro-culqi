package uculqi

import "time"

type Order struct {
	ID        string
	CreatedAt time.Time
	Token     string
	Products  []*Product
	Info      *PersonInfo
	Currency  *Currency
	Card      string
	Discount  float64
	Metadata  map[string]interface{}
}
