// Package service implements application business logic. Each logic group in own file.
package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/repository"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/hasher"
)

// UsersService -.
type UsersService struct {
	repo   repository.Users
	hasher hasher.PasswordHasher
	c      pb.SpawnClient
}

// NewUsersService -.
func NewUsersService(repo repository.Users, client pb.SpawnClient,
	hasher hasher.PasswordHasher) *UsersService {

	return &UsersService{
		repo:   repo,
		c:      client,
		hasher: hasher,
	}
}

// CreateUser -.
func (us *UsersService) CreateUser(ctx context.Context, user entity.User) error {
	response, err := us.c.Generate(ctx, &pb.Empty{})
	if err != nil {
		return fmt.Errorf("UsersService - CreateUser - Generate: %w", err)
	}

	user.Salt = response.Str
	hashed, err := us.hasher.Hash(user.Salt, user.Password)
	if err != nil {
		return fmt.Errorf("UsersService - CreateUser - Hash: %w", err)
	}

	user.Password = hashed

	err = us.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return entity.ErrUserAlreadyExists
		}
		return fmt.Errorf("UsersService - CreateUser - Create: %w", err)
	}

	return nil
}

// GetByEmail -.
func (us *UsersService) GetByEmail(ctx context.Context, email string) (entity.User, error) {

	user, err := us.repo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			return entity.User{}, entity.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("UsersService - GetByEmail: %w", err)
	}

	return user, nil
}
