package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type LoggingService struct {
	next PriceService
}

func NewLoggingService(next PriceService) PriceService {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		log.Info().
			Any("requestID", ctx.Value("requestID")).
			Dur("took", time.Since(begin)).
			AnErr("err", err).
			Str("ticker", ticker).
			Float64("price", price).
			Msg("Fetch price")
	}(time.Now())

	return l.next.FetchPrice(ctx, ticker)
}
