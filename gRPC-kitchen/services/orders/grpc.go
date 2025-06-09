package main

import (
	"log"
	"net"

	handler "github.com/Deirror/gRPC-Kitchen/services/orders/handler/orders"
	"github.com/Deirror/gRPC-Kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewOrdersGrpcHandler(grpcServer, orderService)

	log.Println("Starting gRPC server on:", s.addr)

	return grpcServer.Serve(lis)
}
