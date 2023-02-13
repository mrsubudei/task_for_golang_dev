package mongodb_test

import (
	"context"
	"testing"

	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/go-playground/assert/v2"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/entity"
	m "github.com/mrsubudei/task_for_golang_dev/users-service/internal/repository/mongodb"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setup(t *testing.T, ctx context.Context) (*mongo.Database, *mim.Server) {
	server, err := mim.Start(ctx, "5.0.2")
	if err != nil {
		t.Fatal("error while starting", err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(server.URI()))
	if err != nil {
		t.Fatal("error while connect", err)
	}
	require.NoError(t, err)
	db := client.Database("test")
	collection := db.Collection("users")
	index := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := collection.Indexes().CreateOne(context.Background(), index); err != nil {
		t.Fatal("error while create index", err)
	}
	return db, server
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, server := setup(t, ctx)
	defer server.Stop(ctx)
	repo := m.NewUsersRepo(db)

	user := entity.User{
		Password: "pass",
		Salt:     "salty",
		Email:    "johndoe@example.com",
	}
	t.Run("OK", func(t *testing.T) {

		err := repo.Create(ctx, user)
		require.NoError(t, err)

		var result entity.User
		err = db.Collection("users").FindOne(ctx, bson.M{"email": user.Email}).Decode(&result)
		require.NoError(t, err)
		result.Id = primitive.NilObjectID
		assert.Equal(t, user, result)
	})

	t.Run("Error user exists", func(t *testing.T) {
		err := repo.Create(ctx, user)
		assert.Equal(t, entity.ErrUserAlreadyExists, err)
	})
}

func TestGetByEmail(t *testing.T) {
	ctx := context.Background()
	db, server := setup(t, ctx)
	defer server.Stop(ctx)
	repo := m.NewUsersRepo(db)
	user := entity.User{
		Password: "pass",
		Salt:     "salty",
		Email:    "johndoe@example.com",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatal("Unexpected error", err)
	}

	t.Run("OK", func(t *testing.T) {
		found, err := repo.GetByEmail(ctx, user.Email)
		require.NoError(t, err)
		found.Id = primitive.NilObjectID
		assert.Equal(t, user, found)
	})

	t.Run("Error not found", func(t *testing.T) {
		_, err := repo.GetByEmail(ctx, "notExistingEmail")
		assert.Equal(t, entity.ErrUserNotFound, err)
	})
}
