package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	server "ozon_fintech_test/internal/service/controllers"
	pb "ozon_fintech_test/internal/service/proto/generated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	gatewayAdder string
	gatewatPort  string
	gRPCAddre    string
	gRPCPort     string

	shortUrlController *server.ShortUrlController
}

func MakeServer(controller *server.ShortUrlController, gatewayAdder, gatewatPort, gRPCAddre, gRPCPort string) *Server {
	return &Server{
		gatewayAdder:       gatewayAdder,
		gatewatPort:        gatewatPort,
		gRPCAddre:          gRPCAddre,
		gRPCPort:           gRPCPort,
		shortUrlController: controller,
	}
}

func (server *Server) runRestGateway(httpGateAdr, gRPCAdr string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	error := pb.RegisterShortURLHandlerFromEndpoint(ctx, mux, gRPCAdr, opts)
	if error != nil {
		log.Fatalln("Error to register gRPC server")
	}

	error = http.ListenAndServe(httpGateAdr, mux)
	if error != nil {
		log.Fatalln("Error to start HTTP Gateway")
	}
}

func (server *Server) runGRPCServer(ctx context.Context, gRPCServer *grpc.Server) {
	log.Println("Starting gRPC server...")
	listener, error := net.Listen("tcp", fmt.Sprintf("%s:%s", server.gRPCAddre, server.gRPCPort))
	if error != nil {
		log.Fatalf("Failed to start gRPC at %s:%s", server.gRPCAddre, server.gRPCPort)
	}
	gRPCServer.Serve(listener)
}

func (server *Server) runGatewayServer(ctx context.Context, filledMux *runtime.ServeMux) {
	log.Println("Starting HTTP Gateway..")
	mux := server.getFillGRPCMux(ctx)
	error := http.ListenAndServe(fmt.Sprintf("%s:%s", server.gatewayAdder, server.gatewatPort), mux)
	if error != nil {
		log.Fatalf("Failed to start HTTP Gateway at %s:%s", server.gatewayAdder, server.gatewatPort)
	}
}

func (server *Server) getFillGRPCMux(ctx context.Context) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	error := pb.RegisterShortURLHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", server.gRPCAddre, server.gRPCPort), opts)
	if error != nil {
		log.Fatalln("Failed to register GRPC Service")
	}
	return mux
}

func (server *Server) Run() {
	log.Println("Starting Server...")
	gRPCServer := grpc.NewServer()
	mux := server.getFillGRPCMux(context.TODO())
	pb.RegisterShortURLServer(gRPCServer, server.shortUrlController)
	go server.runGRPCServer(context.TODO(), gRPCServer)
	server.runGatewayServer(context.TODO(), mux)
}
