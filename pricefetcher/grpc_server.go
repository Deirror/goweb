package main

import (
	"context"
	"math/rand/v2"
	"net"
	"pricefetcher/protobuf"

	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceService) error {
	grpcPriceFetcherServer := NewGRPCPriceFetcherServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	protobuf.RegisterPriceFetcherServer(server, grpcPriceFetcherServer)

	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceService
	protobuf.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcherServer(svc PriceService) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *protobuf.PriceRequest) (*protobuf.PriceResponse, error) {
	reqId := rand.IntN(10000)
	ctx = context.WithValue(ctx, "requestID", reqId)

	price, err := s.svc.FetchPrice(ctx, req.Ticker)

	if err != nil {
		return nil, err
	}

	resp := &protobuf.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}

	return resp, nil
}
