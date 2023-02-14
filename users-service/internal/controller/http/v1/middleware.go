package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

func (h *UsersHandler) checkValues(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}
		err := h.parseJson(w, r, &user)
		if err != nil {
			h.l.Error(fmt.Errorf("v1 - checkValues - parseJson: %w", err))
			return
		}

		errorMsg := ErrMessage{
			code:  http.StatusBadRequest,
			Error: EmptyFiledRequest,
		}

		switch {
		case user.Email == "":
			errorMsg.Detail = EmailFieldEmpty
			h.writeResponse(w, errorMsg)
			return
		case user.Password == "":
			errorMsg.Detail = PasswordFieldEmpty
			h.writeResponse(w, errorMsg)
			return
		}

		_, err = mail.ParseAddress(user.Email)
		if err != nil {
			h.writeResponse(w, ErrMessage{
				code:  http.StatusBadRequest,
				Error: WrongEmailFormat,
			})
			return
		}

		ctx := context.WithValue(r.Context(), Key(userKey), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
