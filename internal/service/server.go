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
	gatewayAdder       string
	gatewatPort        string
	gRPCAddre          string
	gRPCPort           string
	httpGateway        *http.Server
	gRPCServer         *grpc.Server
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
	server.httpGateway = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", server.gatewayAdder, server.gatewatPort),
		Handler: mux,
	}
	error := server.httpGateway.ListenAndServe()
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

func (server *Server) Run(ctx context.Context) {
	log.Println("Starting Server...")
	server.gRPCServer = grpc.NewServer()
	mux := server.getFillGRPCMux(context.TODO())
	pb.RegisterShortURLServer(server.gRPCServer, server.shortUrlController)

	go server.runGRPCServer(ctx, server.gRPCServer)
	go server.runGatewayServer(ctx, mux)

	for {
		select {
		case <-ctx.Done():
			server.Shutdown()
			return
		}
	}
}

func (server *Server) Shutdown() {
	log.Println("Shooting down http gateway...")
	server.httpGateway.Shutdown(context.TODO())
	log.Println("Shooting down gRPC server...")
	server.gRPCServer.GracefulStop()
}
