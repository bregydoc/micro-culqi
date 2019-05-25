package uculqi

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestInvoice_ToFlatten(t *testing.T) {
	invoice, err := newInvoiceWithMinimal(&CompanyInfo{
		Name: "TestCompany",
	}, &MinimalInformation{
		Products: []*Product{{Name: "TestProduct", Price: 10.0}},
		Email:    "example@example.com",
		Currency: PEN,
		Token:    "XXX",
	})

	if err != nil {
		t.Error("failed to create new invoice, error: ", err)
		return
	}
	subTotal := 0.0
	for _, p := range invoice.Order.Products {
		subTotal += p.Price
	}
	type fields struct {
		ID         string
		Order      *Order
		Email      string
		Charged    bool
		Company    *CompanyInfo
		ChargedAt  time.Time
		ExternalID string
	}
	type args struct {
		description []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InvoiceFlatten
	}{
		{name: "full correct",
			fields: fields{
				ID:         invoice.ID,
				ExternalID: invoice.ExternalID,
				ChargedAt:  invoice.ChargedAt,
				Company:    invoice.Company,
				Order:      invoice.Order,
				Email:      invoice.Email,
				Charged:    invoice.Charged,
			},
			want: &InvoiceFlatten{
				ID:            invoice.ID,
				Description:   "",
				Date:          invoice.ChargedAt.Format("15:04:05 January 2, 2006"),
				ChargedCard:   invoice.Order.Card,
				Items:         invoice.Order.Products,
				SubTotal:      fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal),
				Discount:      fmt.Sprintf("%s -%0.2f", invoice.Order.Currency.Symbol, invoice.Order.Discount),
				TotalCost:     fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Symbol, subTotal-invoice.Order.Discount),
				CustomerName:  invoice.Order.Info.Name,
				ChargedAmount: fmt.Sprintf("%s %0.2f", invoice.Order.Currency.Code, subTotal-invoice.Order.Discount),
				CompanyName:   invoice.Company.Name,
			},
			args: args{description: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoice := &Invoice{
				ID:         tt.fields.ID,
				Order:      tt.fields.Order,
				Email:      tt.fields.Email,
				Charged:    tt.fields.Charged,
				Company:    tt.fields.Company,
				ChargedAt:  tt.fields.ChargedAt,
				ExternalID: tt.fields.ExternalID,
			}
			if got := invoice.ToFlatten(tt.args.description...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invoice.ToFlatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
