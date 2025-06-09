package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")

	svc := NewLoggingService(&StockFetcher{})

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
