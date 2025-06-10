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
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	var (
		jsonAddr = flag.String("json", ":3000", "listen address the json service is running")
		grpcAddr = flag.String("grpc", ":4000", "listen address the grpc service is running")
		ctx      = context.Background()
		svc      = NewLoggingService(&StockFetcher{})
	)

	go makeGRPCServerAndRun(*grpcAddr, svc)

	go func() {
		time.Sleep(2 * time.Second)
		grpcClient, err := client.NewGRPCClient(*grpcAddr)
		if err != nil {
			fmt.Printf("1%+v", err)
			return
		}

		for {
			time.Sleep(3 * time.Second)
			req := protobuf.PriceRequest{
				Ticker: "AMD",
			}

			resp, err := grpcClient.FetchPrice(ctx, &req)
			if err != nil {
				fmt.Printf("2%+v", err)
				return
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	server := NewJSONAPIServer(*jsonAddr, svc)
	server.Run()
}
