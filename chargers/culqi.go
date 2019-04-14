package chargers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bregydoc/micro-culqi"
	"io/ioutil"
	"net/http"
	"time"
)

const chargeUrl = "https://api.culqi.com/v2/charges"

type CulqiCharger struct {
	MerchantCode string
	ApiKey       string
}

func (q *CulqiCharger) ExecuteCharge(invoice *uculqi.Invoice) (*uculqi.ChargeResponse, error) {
	if !invoice.IsValid() {
		return nil, uculqi.ErrInvalidInvoice
	}

	charge := q.invoiceToCharge(invoice)

	apiKey := q.ApiKey

	data, err := json.Marshal(charge)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, chargeUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("User-Agent", "micro-culqi-go/1.0")

	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusCreated {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		r := new(errorResponseFromCulqi)
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

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cResponse := new(chargeResponse)
	err = json.Unmarshal(data, cResponse)
	if err != nil {
		return nil, err
	}

	r := &uculqi.ChargeResponse{
		ID:          cResponse.ID,
		At:          time.Now(),
		How:         float64(cResponse.Amount) / float64(invoice.Order.Currency.Multiplier),
		UserMessage: cResponse.Outcome.UserMessage,
	}

	return r, err
}
