package v1_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
	v1 "github.com/mrsubudei/task_for_golang_dev/users-service/internal/controller/http/v1"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	mock_service "github.com/mrsubudei/task_for_golang_dev/users-service/internal/service/mock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/logger"
)

type test struct {
	name         string
	path         string
	requsetBody  string
	user         entity.User
	wantStatus   int
	responseBody string
}

func setup(t *testing.T) (*mock_service.UsersMockService, *logger.Logger) {
	cfg, err := config.NewConfig("../../../../config.yml", "../../../../env.example")
	if err != nil {
		t.Fatal(err)
	}

	service := mock_service.NewUsersMockService()
	l := logger.New(cfg.Logger.Level)

	return service, l
}

func Test_createUser(t *testing.T) {
	mockService, l := setup(t)

	tests := []test{
		{
			name: "OK",
			requsetBody: `{
				"email": "sfef111@mail.ru",
				"password": "pass"
			}`,
			wantStatus:   http.StatusCreated,
			responseBody: `{}`,
		},
		{
			name: "Error user already exists",
			requsetBody: `{
				"email": "exist@mail.ru",
				"password": "pass"
			}`,
			wantStatus:   http.StatusConflict,
			responseBody: `{"error":"user with such email already exists"}`,
		},
		{
			name: "Error internal",
			requsetBody: `{
				"email": "internal@error",
				"password": "pass"
			}`,
			wantStatus:   http.StatusInternalServerError,
			responseBody: `{"error":"Internal Server Error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/create-user",
				bytes.NewReader([]byte(tt.requsetBody)))

			mux := chi.NewRouter()
			v1.NewRouter(mux, l, mockService)
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("want: %v, got: %v", tt.wantStatus, rec.Code)
			} else if rec.Body.String() != tt.responseBody {
				t.Fatalf("want: %v, got: %v", tt.responseBody, rec.Body.String())
			}
		})
	}
}

func Test_getByEmail(t *testing.T) {
	mockService, l := setup(t)
	user := entity.User{
		Email:    "sheldon@mar.com",
		Password: "pass",
	}
	if err := mockService.CreateUser(context.Background(), user); err != nil {
		t.Fatal("Unexpected error:", err)
	}

	tests := []test{
		{
			name:         "OK",
			wantStatus:   http.StatusOK,
			path:         "/get-user/sheldon@mar.com",
			responseBody: `{"id":"000000000000000000000000","email":"sheldon@mar.com","password":"pass"}`,
		},
		{
			name:         "Error not found",
			wantStatus:   http.StatusNotFound,
			path:         "/get-user/not@found",
			responseBody: `{"error":"user does not exist"}`,
		},
		{
			name:         "Error internal",
			wantStatus:   http.StatusInternalServerError,
			path:         "/get-user/internal@error",
			responseBody: `{"error":"Internal Server Error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.path,
				bytes.NewReader([]byte(tt.requsetBody)))

			mux := chi.NewRouter()
			v1.NewRouter(mux, l, mockService)
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("want: %v, got: %v", tt.wantStatus, rec.Code)
			} else if rec.Body.String() != tt.responseBody {
				t.Fatalf("want: %v, got: %v", tt.responseBody, rec.Body.String())
			}
		})
	}
}
