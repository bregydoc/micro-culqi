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
		t.Error("error should be different to nil because the service is void")
	}
	if ok.Ok {
		t.Error("'ok' response should be different to true because the service is void")
	}
}
