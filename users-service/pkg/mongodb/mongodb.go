package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Timeout         = 10 * time.Second
	UsersCollection = "users"
	EmailField      = "email"
)

// NewClient established connection to a mongoDb instance using provided URI and auth credentials.
func NewMongoDB(cfg *config.Config) (*mongo.Database, error) {
	opts := options.Client().ApplyURI(cfg.MongoDB.URI)
	if cfg.MongoDB.User != "" && cfg.MongoDB.Password != "" {
		opts.SetAuth(options.Credential{
			Username: cfg.MongoDB.User,
			Password: cfg.MongoDB.Password,
		})
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.MongoDB.Name)
	collection := db.Collection(UsersCollection)
	index := mongo.IndexModel{
		Keys:    bson.M{EmailField: 1},
		Options: options.Index().SetUnique(true),
	}

	if _, err := collection.Indexes().CreateOne(context.Background(), index); err != nil {
		return nil, err
	}

	return db, nil
}

func IsDuplicate(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}

	return false
}
