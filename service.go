package uculqi

import (
	"context"
	"github.com/bregydoc/micro-culqi/proto"
	"io/ioutil"
)

type Service struct {
	Company *Company
}

func (s *Service) CreateNewInvoice(c context.Context, minimal *pculqi.MinimalInvoice) (*pculqi.Invoice, error) {
	products := make([]*Product, 0)
	for _, p := range minimal.Products {
		var currency *Currency
		switch p.Currency.Code {
		case pculqi.PEN.Code:
			currency = PEN
		case pculqi.USD.Code:
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
	switch minimal.Currency.Code {
	case pculqi.PEN.Code:
		currency = PEN
	case pculqi.USD.Code:
		currency = USD
	default:
		currency = PEN
	}
	invoice, err := s.Company.generateNewInvoice(
		minimal.Token,
		minimal.Email,
		products,
		currency,
	)

	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), nil
}

func (s *Service) CreateNewInvoiceWithOrder(c context.Context, order *pculqi.Order) (*pculqi.Invoice, error) {
	invoice, err := s.Company.generateNewInvoiceWithOrder(s.orderProtoToNative(order))
	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), nil
}

func (s *Service) ChargeInvoice(c context.Context, invoiceID *pculqi.InvoiceID) (*pculqi.Invoice, error) {
	invoice, err := s.Company.chargeInvoiceByID(invoiceID.Id)
	if err != nil {
		return nil, err
	}

	return s.invoiceNativeToProto(invoice), nil
}

func (s *Service) SendInvoiceAsEmail(c context.Context, invoiceID *pculqi.InvoiceID) (*pculqi.Invoice, error) {
	invoice, err := s.Company.sendInvoiceAsEmail(invoiceID.Id)
	if err != nil {
		return nil, err
	}
	return s.invoiceNativeToProto(invoice), err
}

func (s *Service) GetInvoiceByID(c context.Context, invoiceID *pculqi.InvoiceID) (*pculqi.Invoice, error) {
	invoice, err := s.Company.Store.GetInvoice(invoiceID.Id)
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
