package v1_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"testing"

	v1 "github.com/mrsubudei/task_for_golang_dev/users-service/internal/controller/http/v1"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

func getMockHandler(t *testing.T) http.HandlerFunc {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(v1.Key("user")).(entity.User)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else if user.Email == "" || user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
	return mockHandler
}

func Test_checkValues(t *testing.T) {

	tests := []struct {
		name        string
		wantStatus  int
		wantResult  string
		requsetBody string
	}{
		{
			name:        "OK",
			requsetBody: `{"email": "sfef111@mail.ru", "password": "pass"}`,
			wantStatus:  http.StatusOK,
		},
		{
			name:        "Error wrong email format",
			requsetBody: `{"email": "wrong", "password": "pass"}`,
			wantStatus:  http.StatusBadRequest,
			wantResult:  `{"error":"email has wrong format"}`,
		},
		{
			name:        "Error empty email",
			requsetBody: `{"password": "pass"}`,
			wantStatus:  http.StatusBadRequest,
			wantResult:  `{"error":"json body has empty fields","detail":"'email:' field is empty"}`,
		},
		{
			name:        "Error empty password",
			requsetBody: `{"email": "bobik@mail.ru"}`,
			wantStatus:  http.StatusBadRequest,
			wantResult:  `{"error":"json body has empty fields","detail":"'password:' field is empty"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := v1.UsersHandler{}
			mockHandler := getMockHandler(t)
			handlerToTest := h.CheckValues(mockHandler)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/create-user",
				bytes.NewReader([]byte(tt.requsetBody)))

			handlerToTest.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("want: %v, got: %v", tt.wantStatus, rec.Code)
			} else if rec.Body.String() != tt.wantResult {
				t.Fatalf("want: %v, got: %v", tt.wantResult, rec.Body.String())
			}
		})
	}
}
