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
)

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

	c.With(h.CheckValues).Post("/create-user", h.createUser)
	c.Get("/get-user/{email}", h.getByEmail)

}

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
