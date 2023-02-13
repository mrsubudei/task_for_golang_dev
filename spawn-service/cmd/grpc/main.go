package main

import (
	"log"

	"github.com/mrsubudei/task_for_golang_dev/spawn-service/internal/config"
	"github.com/mrsubudei/task_for_golang_dev/spawn-service/internal/server"
	"github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto/logger"
)

func main() {
	// config
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// logger
	l := logger.New(cfg.Logger.Level)

	// Grpc server
	if err := server.NewGrpcServer(l).Start(cfg); err != nil {
		l.Error("Failed creating gRPC server", err)
		return
	}
}
