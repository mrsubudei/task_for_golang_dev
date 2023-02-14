// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/service"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/mrsubudei/task_for_golang_dev/users-service/docs"
)

// NewRouter -.
// Swagger spec:
// @title       Users-service
// @version     1.0
// @host        localhost:8081
// @BasePath    /v1
func NewRouter(c *chi.Mux, l logger.Interface, service service.Service) {

	c.Use(middleware.RequestID)
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Use(middleware.Timeout(60 * time.Second))

	h := &UsersHandler{
		service: service,
		l:       l,
		c:       c,
	}
	c.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))
	c.With(h.CheckValues).Post("/create-user", h.createUser)
	c.Get("/get-user/{email}", h.getByEmail)

}

// ParseJson -.
func (h *UsersHandler) ParseJson(w http.ResponseWriter, r *http.Request,
	user *entity.User) error {

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		h.WriteResponse(w, ErrMessage{Code: http.StatusBadRequest,
			Error: JsonNotCorrect})
		return fmt.Errorf(WrongDataFormat)
	}

	return nil
}

// WriteResponse -.
func (h *UsersHandler) WriteResponse(w http.ResponseWriter, ans Answer) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonResp, err := json.Marshal(ans)
	if err != nil {
		h.l.Error(fmt.Errorf("v1 - WriteResponse - Marshal: %w", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(ans.getCode())
	if _, err = w.Write(jsonResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
