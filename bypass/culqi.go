package bypass

import "strings"

type Culqi struct {
	apiKey     string
	apiVersion string
	env        string
}

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
