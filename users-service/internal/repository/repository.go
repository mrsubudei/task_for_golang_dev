package repository

import (
	"context"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
)

type Users interface {
	Create(ctx context.Context, user entity.User) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}
