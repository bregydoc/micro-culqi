package uculqi

import "fmt"

type InvoiceFlatten struct {
	ID            string
	CompanyName   string
	Date          string
	CostumerName  string
	TotalCost     string
	SubTotal      string
	Discount      string
	ChargedCard   string
	ChargedAmount string
	Items         []*Product
}

func (invoice *Invoice) ToFlatten() *InvoiceFlatten {
	subTotal := 0.0
	for _, p := range invoice.Order.Products {
		subTotal += p.Price
	}
	return &InvoiceFlatten{
		ID:            invoice.ID,
		Date:          invoice.ChargedAt.Format("15:04:05 January 2, 2006"),
		ChargedCard:   invoice.Order.Card,
		Items:         invoice.Order.Products,
		SubTotal:      fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal),
		Discount:      fmt.Sprintf("%s -%0.2f", invoice.Order.Currency.Symbol, invoice.Order.Discount),
		TotalCost:     fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal-invoice.Order.Discount),
		CostumerName:  invoice.Order.Info.Name,
		ChargedAmount: fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Code, subTotal),
		CompanyName:   invoice.Company.Name,
	}
}
