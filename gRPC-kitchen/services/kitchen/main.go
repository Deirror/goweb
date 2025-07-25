package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.
			NewCredentials()))

	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}

	return conn
}

func main() {
	s := NewHttpServer(":1000")
	s.Run()
}
