package uculqi

import (
	"fmt"
	"strings"
)

type InvoiceFlatten struct {
	ID            string
	Description   string
	CompanyName   string
	Date          string
	CustomerName  string
	TotalCost     string
	SubTotal      string
	Discount      string
	ChargedCard   string
	ChargedAmount string
	Items         []*Product
}

func (invoice *Invoice) ToFlatten(description ...string) *InvoiceFlatten {
	desc := strings.Join(description, "/n")
	subTotal := 0.0
	for _, p := range invoice.Order.Products {
		subTotal += p.Price
	}
	return &InvoiceFlatten{
		ID:            invoice.ID,
		Description:   desc,
		Date:          invoice.ChargedAt.Format("15:04:05 January 2, 2006"),
		ChargedCard:   invoice.Order.Card,
		Items:         invoice.Order.Products,
		SubTotal:      fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal),
		Discount:      fmt.Sprintf("%s -%0.2f", invoice.Order.Currency.Symbol, invoice.Order.Discount),
		TotalCost:     fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal-invoice.Order.Discount),
		CustomerName:  invoice.Order.Info.Name,
		ChargedAmount: fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Code, subTotal-invoice.Order.Discount),
		CompanyName:   invoice.Company.Name,
	}
}
