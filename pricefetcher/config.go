package main

// Configurations for the business logic of the AV's API
type AlphaVantageConfig struct {
	apiKey string
	url    string
}

func NewAlphaVantageConfig(apiKey, url string) *AlphaVantageConfig {
	return &AlphaVantageConfig{
		apiKey: apiKey,
		url:    url,
	}
}
