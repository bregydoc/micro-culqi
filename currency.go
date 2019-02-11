package culqi

// CurrencyType is a currency type
type CurrencyType string

// PEN define a currency type
const PEN CurrencyType = "PEN"

// USD define a currency type
const USD CurrencyType = "USD"

// BCOIN define a currency type
const BCOIN CurrencyType = "BCOIN"

// Currency is a bombo currency
type Currency struct {
	Type            CurrencyType `json:"type"`
	Symbol          string       `json:"symbol"`
	AssociatedImage string       `json:"associated_image"`
	Multiplier      int          `json:"multiplier"`
}
