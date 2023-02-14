// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
)

const DefaultTime = int(time.Second)

type Server struct {
	httpServer *http.Server
	h          http.Handler
	cfg        *config.Config
}

func NewServer(handler *chi.Mux, cfg *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: cfg.Http.Port,
			ReadTimeout: time.Duration(cfg.Http.DefaultReadTimeout *
				DefaultTime),
			WriteTimeout: time.Duration(cfg.Http.DefaultWriteTimeout *
				DefaultTime),
			Handler: handler,
		},
		h:   handler,
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(s.cfg.Http.DefaultShutdownTimeout*DefaultTime))
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
