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
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AddressCity string `json:"address_city"`
	PhoneNumber int64  `json:"phone_number"`
	CountryCode string `json:"country_code"`
}

type chargeParams struct {
	Token            string           `json:"source_id"`
	Email            string           `json:"email"`
	CurrencyCode     string           `json:"currency_code"`
	Amount           int              `json:"amount"`
	Installments     int              `json:"installments"`
	Description      string           `json:"description"`
	AntifraudDetails antifraudDetails `json:"antifraud_details"`
}

type errorResponseFromCulqi struct {
	Object          string `json:"object"`
	Type            string `json:"type"`
	MerchantMessage string `json:"merchant_message"`
}

func (q *CulqiCharger) invoiceToCharge(i *uculqi.Invoice) *chargeParams {
	description := ""
	totalAmount := 0.0
	for i, p := range i.Order.Products {
		totalAmount += p.Price * float64(p.Currency.Multiplier)
		description += fmt.Sprintf("%d) %s --- %s %.2f\n", i, p.Name, p.Currency.Symbol, p.Price)
	}

	names := strings.Split(i.Order.Info.Name, " ")
	firsName, lastName := names[0], names[1]

	return &chargeParams{
		Token:       i.Order.Token,
		Email:       i.Email,
		Description: description,
		Amount:      int(totalAmount),
		AntifraudDetails: antifraudDetails{
			FirstName:   firsName,
			LastName:    lastName,
			CountryCode: i.Order.Info.CountryCode,
			PhoneNumber: i.Order.Info.Phone,
			Address:     i.Order.Info.Address,
			AddressCity: i.Order.Info.AddressCity,
		},
		CurrencyCode: i.Order.Currency.Code,
	}

}
