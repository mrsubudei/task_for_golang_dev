package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"

	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/repository"
)

type UsersService struct {
	repo repository.Users
	c    pb.SpawnClient
}

func NewUsersService(repo repository.Users, client pb.SpawnClient) *UsersService {
	return &UsersService{
		repo: repo,
		c:    client,
	}
}

func (us *UsersService) CreateUser(ctx context.Context, user entity.User) error {
	response, err := us.c.Generate(ctx, &pb.Empty{})
	if err != nil {
		return fmt.Errorf("UsersService - CreateUser - Generate: %w", err)
	}

	// combine and hash raw password with generated salt string
	h := md5.New()
	io.WriteString(h, response.Str)
	io.WriteString(h, user.Password)
	hashed := fmt.Sprintf("%x", h.Sum(nil))

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
