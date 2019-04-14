package uculqi

import "github.com/bregydoc/micro-culqi/littleid"

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
			ID:       littleid.New(),
			Currency: info.Currency,
			Token:    info.Token,
			Info: &PersonInfo{
				Email: info.Email,
			},
			Products: info.Products,
			Metadata: map[string]interface{}{
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
