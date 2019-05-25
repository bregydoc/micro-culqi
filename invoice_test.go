package uculqi

import (
	"testing"
	"time"
)

func TestInvoice_IsValid(t *testing.T) {
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

	type fields struct {
		ID         string
		Order      *Order
		Email      string
		Charged    bool
		Company    *CompanyInfo
		ChargedAt  time.Time
		ExternalID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "null invoice", fields: fields{}, want: false},
		{name: "semi-null invoice", fields: fields{
			ID:        "",
			Email:     "",
			Company:   &CompanyInfo{},
			Order:     &Order{},
			ChargedAt: time.Now(),
		}, want: false},
		{name: "semi-null invoice", fields: fields{
			ID:        "",
			Email:     "",
			Company:   &CompanyInfo{Name: ""},
			Order:     &Order{Products: []*Product{}},
			ChargedAt: time.Now(),
		}, want: false},
		{name: "correct generated invoice", fields: fields{
			invoice.ID,
			invoice.Order,
			invoice.Email,
			invoice.Charged,
			invoice.Company,
			invoice.ChargedAt,
			invoice.ExternalID,
		}, want: true},
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
			if got := invoice.IsValid(); got != tt.want {
				t.Errorf("Invoice.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
