package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/service"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/logger"
)

type UsersHandler struct {
	service service.Service
	l       logger.Interface
	c       *chi.Mux
}

func (h *UsersHandler) createUser(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(Key(userKey)).(entity.User)
	if !ok {
		h.l.Error(fmt.Errorf("v1 - createUser - TypeAssertion:"+
			"got data of type %T but wanted %T", user, entity.User{}))
		h.WriteResponse(w, ErrMessage{Code: http.StatusInternalServerError,
			Error: http.StatusText(http.StatusInternalServerError)})
		return
	}

	err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			h.WriteResponse(w, ErrMessage{Code: http.StatusConflict,
				Error: entity.ErrUserAlreadyExists.Error()})
			return
		}
		h.l.Error(fmt.Errorf("v1 - createUser - CreateUser: %w", err))
		h.WriteResponse(w, ErrMessage{Code: http.StatusInternalServerError,
			Error: http.StatusText(http.StatusInternalServerError)})
		return
	}

	h.WriteResponse(w, ErrMessage{
		Code: http.StatusCreated,
	})
}

func (h *UsersHandler) getByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := h.service.GetByEmail(r.Context(), email)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			h.WriteResponse(w, ErrMessage{Code: http.StatusNotFound,
				Error: entity.ErrUserNotFound.Error()})
			return
		}
		h.l.Error(fmt.Errorf("v1 - getByEmail - GetByEmail: %w", err))
		h.WriteResponse(w, ErrMessage{Code: http.StatusInternalServerError,
			Error: http.StatusText(http.StatusInternalServerError)})
		return
	}

	respone := Respone{
		Code:     http.StatusOK,
		Id:       user.Id,
		Email:    user.Email,
		Salt:     user.Salt,
		Password: user.Password,
	}

	h.WriteResponse(w, respone)
}
