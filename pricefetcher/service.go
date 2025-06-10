package main

import (
	"context"
	"fmt"
)

// PriceService is an interface which can fetch a stock price.
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// StockFetcher implements the PriceFetcher interface.
type StockFetcher struct{}

func (s *StockFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"AMD":     121.0,
	"NVIDIA":  145.0,
	"DEIRROR": 441.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
