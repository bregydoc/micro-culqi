package uculqi

import (
	"github.com/bregydoc/micro-culqi/proto"
	"github.com/golang/protobuf/ptypes/any"

	"time"
)

func (s *Service) invoiceProtoToNative(invoice *pculqi.Invoice) *Invoice {
	products := make([]*Product, 0)
	for _, p := range invoice.Order.Products {
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
	chargedAt, _ := time.Parse(time.RFC3339, invoice.ChargedAt)
	createdAt, _ := time.Parse(time.RFC3339, invoice.Order.CreatedAt)
	i := &Invoice{
		ID:         invoice.Id,
		ExternalID: invoice.ExternalId,
		Email:      invoice.Email,
		Charged:    invoice.Charged,
		Company: &CompanyInfo{
			Name:         invoice.Company.Name,
			EmailSupport: invoice.Company.EmailSupport,
			PhoneSupport: invoice.Company.PhoneSupport,
		},
		ChargedAt: chargedAt,
		Order: &Order{
			ID:        invoice.Order.Id,
			Token:     invoice.Order.Token,
			Card:      invoice.Order.Card,
			Discount:  invoice.Order.Discount,
			CreatedAt: createdAt,
			Metadata:  map[string]interface{}{},
			Currency: &Currency{
				Name:       invoice.Order.Currency.Name,
				Code:       invoice.Order.Currency.Code,
				Symbol:     invoice.Order.Currency.Symbol,
				Multiplier: int(invoice.Order.Currency.Multiplier),
			},
			Info: &PersonInfo{
				Name:        invoice.Order.Info.Name,
				Email:       invoice.Order.Info.Email,
				IDNumber:    invoice.Order.Info.IdNumber,
				Phone:       invoice.Order.Info.Phone,
				Address:     invoice.Order.Info.Address,
				AddressCity: invoice.Order.Info.AddressCity,
				CountryCode: invoice.Order.Info.CountryCode,
			},
			Products: products,
		},
	}
	return i
}

func (s *Service) invoiceNativeToProto(invoice *Invoice) *pculqi.Invoice {
	resProducts := make([]*pculqi.Product, 0)
	for _, p := range invoice.Order.Products {
		var currency pculqi.AvailableCurrency
		switch p.Currency {
		case PEN:
			currency = pculqi.AvailableCurrency_PEN
		case USD:
			currency = pculqi.AvailableCurrency_USD
		default:
			currency = pculqi.AvailableCurrency_PEN
		}
		resProducts = append(resProducts, &pculqi.Product{
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Currency:    currency,
		})
	}

	i := &pculqi.Invoice{
		Id:         invoice.ID,
		ExternalId: invoice.ExternalID,
		Email:      invoice.Email,
		Charged:    invoice.Charged,
		Company: &pculqi.CompanyInfo{
			Name:         invoice.Company.Name,
			EmailSupport: invoice.Company.EmailSupport,
			PhoneSupport: invoice.Company.PhoneSupport,
		},
		ChargedAt: invoice.ChargedAt.String(),
		Order: &pculqi.Order{
			Id:        invoice.Order.ID,
			Token:     invoice.Order.Token,
			Card:      invoice.Order.Card,
			Discount:  invoice.Order.Discount,
			CreatedAt: invoice.Order.CreatedAt.String(),
			Metadata:  map[string]*any.Any{},
			Currency: &pculqi.Currency{
				Name:       invoice.Order.Currency.Name,
				Code:       invoice.Order.Currency.Code,
				Symbol:     invoice.Order.Currency.Symbol,
				Multiplier: int32(invoice.Order.Currency.Multiplier),
			},
			Info: &pculqi.PersonInfo{
				Name:        invoice.Order.Info.Name,
				Email:       invoice.Order.Info.Email,
				IdNumber:    invoice.Order.Info.IDNumber,
				Phone:       invoice.Order.Info.Phone,
				Address:     invoice.Order.Info.Address,
				AddressCity: invoice.Order.Info.AddressCity,
				CountryCode: invoice.Order.Info.CountryCode,
			},
			Products: resProducts,
		},
	}

	return i
}

func (s *Service) orderProtoToNative(order *pculqi.Order) *Order {
	products := make([]*Product, 0)
	for _, p := range order.Products {
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
	switch order.Currency.Symbol {
	case "PEN":
		currency = PEN
	case "USD":
		currency = USD
	default:
		currency = PEN
	}

	createdAt, _ := time.Parse(time.RFC3339, order.CreatedAt)

	return &Order{
		ID:        order.Id,
		Token:     order.Token,
		Card:      order.Card,
		Discount:  order.Discount,
		CreatedAt: createdAt,
		Metadata:  map[string]interface{}{},
		Products:  products,
		Currency:  currency,
		Info: &PersonInfo{
			Name:        order.Info.Name,
			Email:       order.Info.Email,
			IDNumber:    order.Info.IdNumber,
			Phone:       order.Info.Phone,
			Address:     order.Info.Address,
			AddressCity: order.Info.AddressCity,
			CountryCode: order.Info.CountryCode,
		},
	}

}
