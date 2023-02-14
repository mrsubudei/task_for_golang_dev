package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	"github.com/mrsubudei/task_for_golang_dev/users-service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const usersCollection = "users"

// UsersRepo -.
type UsersRepo struct {
	db *mongo.Collection
}

// NewUsersRepo -.
func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

// Create -.
func (r *UsersRepo) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		if mongodb.IsDuplicate(err) {
			return entity.ErrUserAlreadyExists
		}
		return fmt.Errorf("UsersRepo - Create - InsertOne: %w", err)
	}

	return nil
}

// GetByEmail -.
func (r *UsersRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entity.User{}, entity.ErrUserNotFound
		}

		return entity.User{}, fmt.Errorf("UsersRepo - GetByEmail - FindOne: %w", err)
	}

	return user, nil
}
