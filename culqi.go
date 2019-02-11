package culqi

// Culqi is an struct for culqi
type Culqi struct {
	config *Config
}

// New generate a new culqi config
func New(config *Config) *Culqi {
	return &Culqi{
		config: &Config{
			MerchantCode: config.MerchantCode,
			APIKey:       config.APIKey,
			APIVersion:   config.APIVersion,
		},
	}
}
