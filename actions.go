package culqi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// PENCurrency is a pen money, (sol peruano)
var PENCurrency = &Currency{
	Type:            PEN,
	Symbol:          "s/",
	AssociatedImage: "pen.png",
	Multiplier:      100,
}

// USDCurrency is a dollar money, (eeuu dollar)
var USDCurrency = &Currency{
	Type:            USD,
	Symbol:          "$",
	AssociatedImage: "usd.png",
	Multiplier:      100,
}

// ChargeUserInformation have the user information for our charge
type ChargeUserInformation struct {
	Token       string
	Email       string
	FirstName   string
	LastName    string
	CountryCode string
}

// Response is the expected response from micro culqi
type Response struct {
	ID                   string    `json:"id"`
	At                   time.Time `json:"at"`
	How                  float64   `json:"how"`
	UserMessageFromCulqi string    `json:"user_message_from_culqi"`
}

// MakeCharge make a charge on culqi, with a currency
func (c *Culqi) MakeCharge(uInformation *ChargeUserInformation, how float64, currency *Currency) (*Response, error) {
	a := how * float64(currency.Multiplier)
	amount := int(a)

	if amount < currency.Multiplier {
		return nil, ErrInvalidProductName
	}

	productName := c.config.ProductName + "_" + strconv.Itoa(amount)

	if len(productName) < 3 {
		return nil, ErrInvalidProductName
	}

	chargeParams := &ChargeParams{
		TokenID:            uInformation.Token,
		Email:              uInformation.Email,
		CountryCode:        uInformation.CountryCode,
		CurrencyCode:       string(currency.Type),
		ProductDescription: productName,
		Amount:             amount,
		FirstName:          uInformation.FirstName,
		LastName:           uInformation.LastName,
	}

	resp, err := c.createCharge(chargeParams)

	if resp.StatusCode != http.StatusCreated {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		r := new(ErrorResponseFromCulqi)
		err = json.Unmarshal(data, r)
		if err != nil {
			return nil, err
		}
		mError := "error from culqi: "
		if r.Object == "error" {
			mError += r.MerchantMessage
		}
		return nil, errors.New(mError)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cResponse := new(ChargeResponse)
	err = json.Unmarshal(data, cResponse)
	if err != nil {
		return nil, err
	}

	r := &Response{
		ID:                   cResponse.ID,
		At:                   time.Now(),
		How:                  float64(cResponse.Amount) / float64(currency.Multiplier),
		UserMessageFromCulqi: cResponse.Outcome.UserMessage,
	}

	return r, nil
}
