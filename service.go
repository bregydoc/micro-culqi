package culqi

import (
	"context"

	"github.com/bregydoc/micro-culqi/grpc"
	"github.com/golang/protobuf/ptypes"
)

// MCulqiService implements a culqi server
type MCulqiService struct {
	c *Culqi
}

func NewMCulqiService(c *Culqi) *MCulqiService {
	return &MCulqiService{
		c: c,
	}
}

// Charge implements the mculqi server
func (m *MCulqiService) Charge(c context.Context, params *mculqi.ChargeParams) (*mculqi.ChargeResponse, error) {
	countryCode := params.UserInfo.CountryCode
	if countryCode == "" {
		countryCode = "PE"
	}
	user := &ChargeUserInformation{
		Email:       params.UserInfo.Email,
		CountryCode: countryCode,
		FirstName:   params.UserInfo.FirstName,
		LastName:    params.UserInfo.LastName,
		Token:       params.UserInfo.Token,
	}
	currency := &Currency{}

	if params.Currency == mculqi.Currency_PEN {
		currency = PENCurrency
	} else if params.Currency == mculqi.Currency_USD {
		currency = USDCurrency
	}

	resp, err := m.c.MakeCharge(user, float64(params.Amount), currency)
	if err != nil {
		return nil, err
	}

	t, err := ptypes.TimestampProto(resp.At)
	if err != nil {
		return nil, err
	}

	return &mculqi.ChargeResponse{
		Id:                   resp.ID,
		At:                   t,
		How:                  float32(resp.How),
		UserMessageFromCulqi: resp.UserMessageFromCulqi,
	}, nil
}
