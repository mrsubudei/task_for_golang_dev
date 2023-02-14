// package mock_service mocks user service
package mock_service

import (
	"context"
	"errors"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

// UsersMockService is a mock for Service interface
type UsersMockService struct {
	Users []entity.User
}

// Creates new UsersMockService
func NewUsersMockService() *UsersMockService {
	return &UsersMockService{}
}

// CreateUser mocks base method
func (um *UsersMockService) CreateUser(ctx context.Context, user entity.User) error {
	if user.Email == "exist@mail.ru" {
		return entity.ErrUserAlreadyExists
	} else if user.Email == "internal@error" {
		return errors.New("internal error")
	}

	um.Users = append(um.Users, user)
	return nil
}

// GetByEmail mocks base method
func (um *UsersMockService) GetByEmail(ctx context.Context,
	email string) (entity.User, error) {

	if email == "not@found" {
		return entity.User{}, entity.ErrUserNotFound
	} else if email == "internal@error" {
		return entity.User{}, errors.New("internal error")
	}

	for _, v := range um.Users {
		if email == v.Email {
			return v, nil
		}
	}
	return entity.User{}, entity.ErrUserNotFound
}
