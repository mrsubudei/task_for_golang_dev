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

// UsersHandler -.
type UsersHandler struct {
	service service.Service
	l       logger.Interface
	c       *chi.Mux
}

// @Summary     Create user
// @Description Creates new user, adds to database
// @ID          create-tser
// @Accept	    json
// @Tags  	    create
// @Accept      json
// @Success     201
// @Failure     500 {object} ErrMessage "Internal server error"
// @Failure     409 {object} ErrMessage "User already exists"
// @Router      /create-user [post]
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

// @Summary     Get user
// @Description Gets user from database using given email
// @ID          get-by-email
// @Produce		json
// @Param		some_email	path		string				true	"email"
// @Tags  	    get
// @Success     200 {object} Response "Ok"
// @Success     404 {object} ErrMessage "Not found"
// @Failure     500 {object} ErrMessage "Internal server error"
// @Router      /get-user{email} [get]
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

	respone := Response{
		Code:     http.StatusOK,
		Id:       user.Id,
		Email:    user.Email,
		Salt:     user.Salt,
		Password: user.Password,
	}

	h.WriteResponse(w, respone)
}
