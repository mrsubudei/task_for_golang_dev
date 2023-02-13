package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	mock_repository "github.com/mrsubudei/task_for_golang_dev/users-service/internal/repository/mock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/service"
	mock_spawn_api "github.com/mrsubudei/task_for_golang_dev/users-service/pkg/api-spawn/mock"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/hasher"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("test: internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
	user entity.User
}

func mockUserService(t *testing.T) (*service.UsersService, *mock_repository.MockUsersRepo) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	hasher := hasher.NewMd5Hasher()
	repo := mock_repository.NewMockUsersRepo(mockCtl)
	spawnClient := mock_spawn_api.NewMockSpawnApi()
	usersService := service.NewUsersService(repo, spawnClient, hasher)

	return usersService, repo
}

func TestCreateDoctor(t *testing.T) {
	t.Parallel()

	eventsService, repo := mockUserService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			user: entity.User{
				Email:    "DuhastVechislavovich@mail.ru",
				Password: "pass",
				Salt:     "sol",
			},
			err: nil,
		},
		{
			name: "Error user already exists",
			mock: func() {
				repo.EXPECT().Create(ctx, gomock.Any()).
					Return(entity.ErrUserAlreadyExists)
			},
			user: entity.User{
				Email:    "DuhastVechislavovich@mail.ru",
				Password: "pass",
				Salt:     "sol",
			},
			err: entity.ErrUserAlreadyExists,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := eventsService.CreateUser(ctx, tc.user)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestGetByEmail(t *testing.T) {
	t.Parallel()

	eventsService, repo := mockUserService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().GetByEmail(ctx, gomock.Any()).Return(entity.User{}, nil)
			},
			res: entity.User{},
			err: nil,
		},
		{
			name: "Error not found",
			mock: func() {
				repo.EXPECT().GetByEmail(ctx, gomock.Any()).Return(entity.User{}, entity.ErrUserNotFound)
			},
			res: entity.User{},
			err: entity.ErrUserNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := eventsService.GetByEmail(ctx, "email")
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}
