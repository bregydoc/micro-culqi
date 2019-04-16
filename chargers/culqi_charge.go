package chargers

import (
	"fmt"
	"github.com/bregydoc/micro-culqi"
	"strings"
)

type chargeResponse struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Outcome struct {
		UserMessage string `json:"user_message"`
	} `json:"outcome"`
}

type antifraudDetails struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty"`
	AddressCity string `json:"address_city,omitempty"`
	PhoneNumber int64  `json:"phone_number,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type chargeParams struct {
	Token            string           `json:"source_id"`
	Email            string           `json:"email,omitempty"`
	CurrencyCode     string           `json:"currency_code,omitempty"`
	Amount           int              `json:"amount,omitempty"`
	Installments     int              `json:"installments,omitempty"`
	Description      string           `json:"description,omitempty"`
	AntifraudDetails antifraudDetails `json:"antifraud_details,omitempty"`
}

type errorResponseFromCulqi struct {
	Object          string `json:"object,omitempty"`
	Type            string `json:"type,omitempty"`
	MerchantMessage string `json:"merchant_message,omitempty"`
}

func (q *CulqiCharger) invoiceToCharge(i *uculqi.Invoice) *chargeParams {
	description := ""
	totalAmount := 0.0
	for i, p := range i.Order.Products {
		totalAmount += p.Price * float64(p.Currency.Multiplier)
		description += fmt.Sprintf("%d) %s --- %s %.2f\n", i, p.Name, p.Currency.Symbol, p.Price)
	}

	names := strings.Split(i.Order.Info.Name, " ")
	var firstName, lastName string

	if len(names) > 1 {
		firstName, lastName = names[0], names[1]
	} else if len(names) > 0 {
		firstName = names[0]
	}

	return &chargeParams{
		Token:       i.Order.Token,
		Email:       i.Email,
		Description: description,
		Amount:      int(totalAmount),
		AntifraudDetails: antifraudDetails{
			FirstName:   firstName,
			LastName:    lastName,
			CountryCode: i.Order.Info.CountryCode,
			PhoneNumber: i.Order.Info.Phone,
			Address:     i.Order.Info.Address,
			AddressCity: i.Order.Info.AddressCity,
		},
		CurrencyCode: i.Order.Currency.Code,
	}

}
