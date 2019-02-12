package culqi

// Culqi is an struct for culqi
type Culqi struct {
	config *Config
}

// New generate a new culqi config
func New(merchantCode, apiKey, apiVersion, productName string) *Culqi {
	return &Culqi{
		config: &Config{
			MerchantCode: merchantCode,
			APIKey:       apiKey,
			APIVersion:   apiVersion,
			ProductName:  productName,
		},
	}
}
