package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
	v1 "github.com/mrsubudei/task_for_golang_dev/users-service/internal/controller/http/v1"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/service"
	api_spawn "github.com/mrsubudei/task_for_golang_dev/users-service/pkg/api-spawn"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/hasher"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/httpserver"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/logger"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/mongodb"

	m "github.com/mrsubudei/task_for_golang_dev/users-service/internal/repository/mongodb"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Logger.Level)

	// MongoDB
	db, err := mongodb.NewMongoDB(cfg)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - NewMongoDB: %w", err))
	}

	// Repository
	repo := m.NewUsersRepo(db)

	// Hasher
	hasher := hasher.NewMd5Hasher()

	// Spawn api client
	spawnClient, err := api_spawn.NewClient(cfg)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - NewClient: %w", err))
	}

	// Service
	service := service.NewUsersService(repo, spawnClient, hasher)

	// Handler
	handler := chi.NewRouter()
	v1.NewRouter(handler, l, service)

	// Http Server
	httpServer := httpserver.NewServer(handler, cfg)

	go func() {
		if err := httpServer.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	fmt.Printf("Server started at http://%s%s\n", cfg.Http.Host, cfg.Http.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
