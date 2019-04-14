package uculqi

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	nilInvoice := Invoice{}
	if nilInvoice.IsValid() {
		t.Error("nilInvoice should be a invalid invoice")
	}

	nilInvoice.ID = ""
	nilInvoice.Email = ""
	nilInvoice.Company = &CompanyInfo{}
	nilInvoice.Order = &Order{}
	nilInvoice.ChargedAt = time.Now()
	if nilInvoice.IsValid() {
		t.Error("nilInvoice should be a invalid invoice")
	}

	nilInvoice.Company.Name = ""
	nilInvoice.Order.Products = []*Product{}
	if nilInvoice.IsValid() {
		t.Error("nilInvoice should be a invalid invoice")
	}

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

	if !invoice.IsValid() {
		t.Error("invoice should be a complete valid invoice")
	}

}

func TestToFlatten(t *testing.T) {
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

	if invoice.ToFlatten().ID == "" {
		t.Error("invoice without id after flatten process")
	}

	if len(invoice.Order.Products) != 1 {
		t.Error("corrupted products after flatten process")
	}
}
