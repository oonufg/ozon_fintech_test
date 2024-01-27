package server

import (
	"ozon_fintech_test/internal/persistance"
)

type Server struct {
	repository persistence.UrlRepository
}

func MakeServer(repository persistence.UrlRepository) *Server {
	return &Server{
		repository: repository,
	}
}

func RunRestGateway() {

}
