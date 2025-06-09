package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type LoggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		log.Info().
			Dur("took", time.Since(begin)).
			AnErr("err", err).
			Float64("price", price).
			Msg("Fetch price")
	}(time.Now())

	return l.next.FetchPrice(ctx, ticker)
}
