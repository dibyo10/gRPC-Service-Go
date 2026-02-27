package main

import (
	"fmt"
	"net"

	handler "github.com/dibyochakraborty/kitchen/services/orders/handler/orders"
	"github.com/dibyochakraborty/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer,orderService)

	fmt.Println("Starting gRPC server on : ", s.addr)

	return grpcServer.Serve(lis)
}
