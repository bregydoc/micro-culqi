package culqi

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// MCulqiConfig configure the micro culqi service
type MCulqiConfig struct {
	Version      string `yaml:"mculqi_version"`
	MerchantCode string `yaml:"merchant_code"`
	APIKey       string `yaml:"api_key"`
	ProductName  string `yaml:"product_name"`
}

// NewMCulqiConfig returns a new config for micro culqi
func NewMCulqiConfig(from string) (*MCulqiConfig, error) {
	data, err := ioutil.ReadFile(from)
	if err != nil {
		return nil, err
	}

	config := new(MCulqiConfig)

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
