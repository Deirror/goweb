package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// PriceService is an interface which can fetch a stock price.
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// MockFetcher implements the PriceService interface.
type MockPriceFetcher struct{}

func (s *MockPriceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return mockFetchPrice(ticker)
}

// Mocks for price fetching
var priceMocks = map[string]float64{
	"AMD":     121.0,
	"NVIDIA":  145.0,
	"DEIRROR": 441.0,
}

func mockFetchPrice(ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}

// Alpha Vantage - Third-party system's API for fetching stock prices
// NOTE: 25 requests per day only for free tier
type AlphaVantagePriceFetcher struct {
	cfg *AlphaVantageConfig
}

func (pf *AlphaVantagePriceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	queryUrl := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", pf.cfg.url, ticker, pf.cfg.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", queryUrl, nil)
	if err != nil {
		return 0.0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return 0.0, fmt.Errorf("request timed out: %w", ctx.Err())
		}
		return 0.0, err
	}
	defer resp.Body.Close()

	var priceResp struct {
		GlobalQuote struct {
			Price string `json:"05. price"`
		} `json:"Global Quote"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return 0.0, err
	}

	fmt.Printf("%+v\n", resp.Body)

	priceStr := priceResp.GlobalQuote.Price
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return price, err
	}

	return price, nil
}
