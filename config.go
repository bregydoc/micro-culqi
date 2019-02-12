package culqi

// Config is a culqi config struct
type Config struct {
	MerchantCode string `json:"merchant_code"`
	APIKey       string `json:"api_key"`
	APIVersion   string `json:"api_version"`
	ProductName  string `json:"product_name"`
}
