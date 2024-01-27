package main

import (
	"net"
	"ozon_fintech_test/internal/persistence"
	server "ozon_fintech_test/internal/service"
	pb "ozon_fintech_test/internal/service/proto/generated"

	"google.golang.org/grpc"
)

func main() {
	gRPCServer := grpc.NewServer()
	go func() {
		listener, _ := net.Listen("tcp", "0.0.0.0:8083")
		gRPCServer.Serve(listener)
	}()

	server := server.MakeServer(persistence.MakeInmemoryRepository())
	pb.RegisterShortURLServer(gRPCServer, server)
	server.RunRestGateway(":8080", "0.0.0.0:8083")

}

//Graceful shutdown сделаю потом
