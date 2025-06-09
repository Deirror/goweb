package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"pricefetcher/client"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")

	svc := NewLoggingService(&StockFetcher{})

	server := NewJSONAPIServer(*listenAddr, svc)
	go server.Run()

	time.Sleep(1 * time.Second)

	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "AMD")
	if err != nil {
		// log here
		return
	}

	fmt.Printf("%+v\n", price)
}
