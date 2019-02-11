package culqi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const chargesBase = "charges"

// ChargeResponse will be returned when the serve mades a charge to culqi servers
type ChargeResponse struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Outcome struct {
		UserMessage string `json:"user_message"`
	} `json:"outcome"`
}

// ChargeParams is a charge struct
type ChargeParams struct {
	TokenID            string `json:"source_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	AddressCity        string `json:"address_city"`
	PhoneNumber        int    `json:"phone_number"`
	CountryCode        string `json:"country_code"`
	CurrencyCode       string `json:"currency_code"`
	Amount             int    `json:"amount"`
	Installments       int    `json:"installments"`
	ProductDescription string `json:"product_description"`
}

func (c *Culqi) createCharge(params *ChargeParams) (*http.Response, error) {

	if c.config.APIVersion == "" {
		c.config.APIVersion = "v2"
	}

	completeURL := defaultBaseURL + c.config.APIVersion + "/" + chargesBase
	apiKey := c.config.APIKey

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, completeURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)

	return resp, err
}
