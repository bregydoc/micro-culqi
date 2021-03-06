package uculqi

import (
	"context"
	"github.com/bregydoc/micro-culqi/proto"
	"testing"
)

func TestUpdateEmailTemplate(t *testing.T) {
	s := &Service{
		Company: &Company{},
	}

	ok, err := s.UpdateEmailTemplate(context.Background(), &pculqi.TemplateData{})
	if err == nil {
		t.Error("error should be different to nil because the service is avoid")
		return
	}

	if ok.Ok {
		t.Error("'ok' response should be different to true because the service is void")
	}
}

func TestCreateNewInvoice(t *testing.T) {
	s := &Service{
		Company: &Company{},
	}

	invoice, err := s.CreateNewInvoice(context.Background(), &pculqi.MinimalInvoice{
		Email: "",
		Products: []*pculqi.Product{
			{Name: "A", Price: 1.2, Currency: pculqi.PEN},
		},
		Currency: pculqi.PEN,
		Token:    "<Neh>",
	})

	if err == nil {
		t.Error("error should be different to nil because the service is avoid")
		return
	}

	if invoice.Id != "" {
		t.Error("error should be different to nil because the service is avoid")
	}
}
