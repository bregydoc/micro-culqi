package bypass

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
)

const baseTokensURL = "https://secure.culqi.com/tokens/"

type tokenPayload struct {
	PublicKey       string                 `json:"public_key"`
	Email           string                 `json:"email"`
	CardNumber      string                 `json:"card_number"`
	CVV             string                 `json:"cvv"`
	ExpirationYear  int                    `json:"expiration_year"`
	ExpirationMonth int                    `json:"expiration_month"`
	Fingerprint     string                 `json:"fingerprint"`
	Metadata        map[string]interface{} `json:"metadata"`
}

// Token ...
type Token struct {
	Value string `json:"id"`
}

type culqiError struct {
	Object  string `json:"object"`
	Message string `json:"merchant_message"`
}

// Details ...
type Details struct {
	Email      string `json:"email"`
	CardNumber string `json:"card_number"`
	ExpYear    int    `json:"expiration_year"`
	ExpMonth   int    `json:"expiration_month"`
	CVV        string `json:"cvv"`
}

func (c *Culqi) getNewToken(payload tokenPayload, sessionID string) (*Token, error) {
	buff := bytes.NewBufferString("")
	err := json.NewEncoder(buff).Encode(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", baseTokensURL, buff)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("X-API-VERSION", c.apiVersion)
	req.Header.Set("X-CULQI-ENV", c.env)
	req.Header.Set("X-CULQI-SESS", sessionID)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		returnedError := new(culqiError)
		err = json.NewDecoder(res.Body).Decode(returnedError)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(returnedError.Message)
	}

	token := new(Token)
	err = json.NewDecoder(res.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil

}

// GenerateNewToken returns a new token fron culqi
func (c *Culqi) GenerateNewToken(cardDetails *Details) (*Token, error) {
	sess, err := c.getNewSessionID()
	if err != nil {
		return nil, err
	}

	return c.getNewToken(tokenPayload{
		Email:           cardDetails.Email,
		CardNumber:      cardDetails.CardNumber,
		CVV:             cardDetails.CVV,
		ExpirationYear:  cardDetails.ExpYear,
		ExpirationMonth: cardDetails.ExpMonth,
		Fingerprint:     "",
		Metadata:        map[string]interface{}{"installments": ""},
	}, sess)
}
