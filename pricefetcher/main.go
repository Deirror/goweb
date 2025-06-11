package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"pricefetcher/client"
	"pricefetcher/protobuf"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func main() {
	// Load env vars for calling third-party price fetcher APIs
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %v", err)
		return
	}

	// Alpha Vantage secrets
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	url := os.Getenv("ALPHA_VANTAGE_URL")

	avSvc := NewLoggingService(&AlphaVantagePriceFetcher{
		cfg: NewAlphaVantageConfig(apiKey, url),
	})

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	var (
		jsonAddr = flag.String("json", ":3000", "listen address the json service is running")
		grpcAddr = flag.String("grpc", ":4000", "listen address the grpc service is running")
		ctx      = context.Background()
		mockSvc  = NewLoggingService(&MockPriceFetcher{})
	)

	go makeGRPCServerAndRun(*grpcAddr, avSvc)

	go func() {
		time.Sleep(1 * time.Second)
		grpcClient, err := client.NewGRPCClient(*grpcAddr)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}

		for {
			time.Sleep(2 * time.Second)
			req := protobuf.PriceRequest{
				Ticker: "AMD",
			}

			resp, err := grpcClient.FetchPrice(ctx, &req)
			if err != nil {
				continue
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	server := NewJSONAPIServer(*jsonAddr, mockSvc)
	server.Run()
}
