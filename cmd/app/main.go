package main

import (
	"ozon_fintech_test/cfg"
	"ozon_fintech_test/internal/persistence"
	server "ozon_fintech_test/internal/service"
	controllers "ozon_fintech_test/internal/service/controllers"
)

func main() {
	cfg := cfg.LoadCFG()
	//ctx := context.Background()
	controller := controllers.MakeShortUrlController(persistence.RepositoryFactoryMethod(cfg))
	server := server.MakeServer(
		controller,
		cfg.HTTP_GATEWAY_ADDRE,
		cfg.HTTP_GATEWAY_PORT,
		cfg.GRPC_ADDR,
		cfg.GRPC_PORT,
	)
	server.Run()
}
