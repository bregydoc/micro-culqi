package bypass

import "strings"

// Culqi is a simple struct to describe your culqi instance
type Culqi struct {
	apiKey     string
	apiVersion string
	env        string
}

// NewCulqi returns a new culqi instance from your culqi api-key
func NewCulqi(apiKey string) *Culqi {
	env := "prod"
	if strings.Contains(apiKey, "test") {
		env = "test"
	}
	return &Culqi{
		apiKey:     apiKey,
		apiVersion: "2",
		env:        env,
	}
}
