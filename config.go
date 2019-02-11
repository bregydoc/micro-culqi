package culqi

// Config is a culqi config struct
type Config struct {
	MerchantCode string `json:"merchant_code"`
	APIKey       string `json:"api_key"`
	APIVersion   string `json:"api_version"`
	BaseURL      string `json:"base_url"`
	UserAgent    string `json:"user_agent"`

	ProductName string `json:"product_name"`
}
