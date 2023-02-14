package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/service"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/logger"
)

type Handler struct {
	service *service.UsersService
	cfg     *config.Config
	l       *logger.Logger
	mux     *chi.Mux
}

func NewHandler(service *service.UsersService, cfg *config.Config,
	logger *logger.Logger) *Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	return &Handler{
		service: service,
		cfg:     cfg,
		l:       logger,
		mux:     r,
	}
}

func (h *Handler) CreateRoutes() {
	h.mux.Get("/create-user", h.CreateUser)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("hi")))
}
