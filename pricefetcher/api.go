package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"pricefetcher/client/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(addr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: addr,
		svc:        svc,
	}
}

func makeHTTPHandlerFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.IntN(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	if ticker == "" {
		return fmt.Errorf("ticker is not set in the url")
	}

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Ticker: ticker,
		Price:  price,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
