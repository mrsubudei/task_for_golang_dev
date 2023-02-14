package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

func (h *UsersHandler) CheckValues(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := entity.User{}
		err := h.ParseJson(w, r, &user)
		if err != nil {
			h.l.Error(fmt.Errorf("v1 - CheckValues - ParseJson: %w", err))
			return
		}

		errorMsg := ErrMessage{
			Code:  http.StatusBadRequest,
			Error: EmptyFiledRequest,
		}

		switch {
		case user.Email == "":
			errorMsg.Detail = EmailFieldEmpty
			h.WriteResponse(w, errorMsg)
			return
		case user.Password == "":
			errorMsg.Detail = PasswordFieldEmpty
			h.WriteResponse(w, errorMsg)
			return
		}

		_, err = mail.ParseAddress(user.Email)
		if err != nil {
			h.WriteResponse(w, ErrMessage{
				Code:  http.StatusBadRequest,
				Error: WrongEmailFormat,
			})
			return
		}

		ctx := context.WithValue(r.Context(), Key(userKey), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
