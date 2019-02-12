package culqi

import "errors"

// ErrInvalidCulqiResponse is an error
var ErrInvalidCulqiResponse = errors.New("invalid culqi response")

// ErrInvalidFirstName is an error
var ErrInvalidFirstName = errors.New("invalid first name")

// ErrInvalidLastName is an error
var ErrInvalidLastName = errors.New("invalid last name")

// ErrInvalidProductName is an error
var ErrInvalidProductName = errors.New("invalid product name")

// ErrorResponseFromCulqi is a generic error culqi response
type ErrorResponseFromCulqi struct {
	Object          string `json:"object"`
	Type            string `json:"type"`
	MerchantMessage string `json:"merchant_message"`
}
