package uculqi

import (
	"github.com/bregydoc/micro-culqi/littleid"
	"time"
)

type MinimalInformation struct {
	Products []*Product
	Currency *Currency
	Email    string
	Token    string
}

func newInvoiceWithMinimal(companyInfo *CompanyInfo, info *MinimalInformation) (*Invoice, error) {
	invoice := &Invoice{
		ID:      littleid.New(),
		Email:   info.Email,
		Company: companyInfo,
		Order: &Order{
			ID:        littleid.New(),
			CreatedAt: time.Now(),
			Currency:  info.Currency,
			Token:     info.Token,
			Info: &PersonInfo{
				Email: info.Email,
			},
			Products: info.Products,
			Metadata: map[string]string{
				"created_with": "minimal",
			},
		},
		Charged: false,
	}

	if !invoice.IsValid() {
		return nil, ErrInvalidInvoice
	}

	return invoice, nil
}

func newFullInvoice(companyInfo *CompanyInfo, order *Order) (*Invoice, error) {
	if order.ID == "" {
		order.ID = littleid.New()
	}

	invoice := &Invoice{
		ID:      littleid.New(),
		Order:   order,
		Email:   order.Info.Email,
		Charged: false,
		Company: companyInfo,
	}

	if !invoice.IsValid() {
		return nil, ErrInvalidInvoice
	}

	return invoice, nil

}
