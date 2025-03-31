package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, mongoConfig Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoConfig.URI)
	clientOptions.SetAuth(options.Credential{
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Connect")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "client.Ping")
	}

	return client, nil
}
