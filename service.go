package uculqi

import (
	"context"
	"github.com/bregydoc/micro-culqi/proto"
	"io/ioutil"
)

type Service struct {
	Company *Company
}

func (s *Service) ExecuteCharge(c context.Context, params *pculqi.MinimalInvoice) (*pculqi.Invoice, error) {
	products := make([]*Product, 0)
	for _, p := range params.Products {
		var currency *Currency
		switch p.Currency {
		case pculqi.AvailableCurrency_PEN:
			currency = PEN
		case pculqi.AvailableCurrency_USD:
			currency = USD
		default:
			currency = PEN
		}
		products = append(products, &Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Currency:    currency,
		})
	}
	var currency *Currency
	switch params.Currency {
	case pculqi.AvailableCurrency_PEN:
		currency = PEN
	case pculqi.AvailableCurrency_USD:
		currency = USD
	default:
		currency = PEN
	}
	invoice, err := s.Company.generateNewInvoice(
		params.Token,
		params.Email,
		products,
		currency,
	)

	if err != nil {
		return nil, err
	}

	invoice, err = s.Company.chargeInvoiceByID(invoice.ID)
	if err != nil {
		return nil, err
	}

	err = s.Company.sendInvoiceAsEmail(invoice.ID)
	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), err
}

func (s *Service) GetInvoiceByID(c context.Context, params *pculqi.InvoiceID) (*pculqi.Invoice, error) {
	invoice, err := s.Company.Store.GetInvoice(params.Id)
	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), nil
}

func (s *Service) UpdateEmailTemplate(c context.Context, params *pculqi.TemplateData) (*pculqi.IsOk, error) {
	err := ioutil.WriteFile(s.Company.TemplateFilename, []byte(params.Data), 0644)
	if err != nil {
		return &pculqi.IsOk{Ok: false}, err
	}

	return &pculqi.IsOk{Ok: true}, nil
}

func (s *Service) ExecuteChargeWithOrder(c context.Context, order *pculqi.Order) (*pculqi.Invoice, error) {
	invoice, err := s.Company.generateNewInvoiceWithOrder(s.orderProtoToNative(order))
	if err != nil {
		return nil, err
	}

	invoice, err = s.Company.chargeInvoiceByID(invoice.ID)
	if err != nil {
		return nil, err
	}

	err = s.Company.sendInvoiceAsEmail(invoice.ID)
	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), err
}
