package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	svc := NewLoggingService(&StockFetcher{})

	price, err := svc.FetchPrice(context.Background(), "AMD")
	if err != nil {
		log.Fatal().Err(err).Msg("fetch failed")
		return
	}

	fmt.Print(price)
}
