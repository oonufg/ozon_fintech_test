package server

import (
	"context"
	"log"
	"net/http"
	"ozon_fintech_test/internal/persistence"
	pb "ozon_fintech_test/internal/service/proto/generated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	repository persistence.UrlRepository
	pb.UnimplementedShortURLServer
}

func MakeServer(repository persistence.UrlRepository) *Server {
	return &Server{
		repository: repository,
	}
}

func (server *Server) RunRestGateway(httpGateAdr, gRPCAdr string) {
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
