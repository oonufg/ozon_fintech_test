package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"ozon_fintech_test/cfg"
	"ozon_fintech_test/internal/persistence"
	server "ozon_fintech_test/internal/service"
	controllers "ozon_fintech_test/internal/service/controllers"
	"syscall"
)

func main() {
	cfg := cfg.LoadCFG()
	ctx := context.Background()
	serverContext, serverWorkerCancel := context.WithCancel(ctx)
	go RunServerWorker(serverContext, cfg)

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	select {
	case sig := <-osSigChan:
		log.Println("Exit signal: ", sig)
		serverWorkerCancel()

	}
}

func RunServerWorker(ctx context.Context, cfg *cfg.Cfg) {
	controller := controllers.MakeShortUrlController(persistence.RepositoryFactoryMethod(cfg))
	server := server.MakeServer(
		controller,
		cfg.HTTP_GATEWAY_ADDRE,
		cfg.HTTP_GATEWAY_PORT,
		cfg.GRPC_ADDR,
		cfg.GRPC_PORT,
	)
	server.Run(ctx)
}
