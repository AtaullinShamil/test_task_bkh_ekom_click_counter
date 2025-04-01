package repository

import (
	"context"
	"time"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	return &MongoRepository{
		db: client.Database("click_tracker"),
	}
}

func (m *MongoRepository) Increase(ctx context.Context, bannerID int) error {
	ts := time.Now().Truncate(time.Minute)

	collection := m.db.Collection("clicks")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"timestamp": ts, "banner_id": bannerID},
		bson.M{"$inc": bson.M{"count": 1}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return errors.Wrap(err, "UpdateOne")
	}

	return nil
}

func (m *MongoRepository) GetStatsBeforeTime(ctx context.Context, ts time.Time) ([]counter.Stat, error) {
	collection := m.db.Collection("clicks")

	filter := bson.M{"timestamp": bson.M{"$lt": ts}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "Find")
	}
	defer cursor.Close(ctx)

	var stats []counter.Stat
	if err := cursor.All(ctx, &stats); err != nil {
		return nil, errors.Wrap(err, "cursor.All")
	}

	return stats, nil
}

func (m *MongoRepository) DeleteStatsBeforeTime(ctx context.Context, ts time.Time) error {
	collection := m.db.Collection("clicks")

	filter := bson.M{"timestamp": bson.M{"$lt": ts}}

	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		errors.Wrap(err, "DeleteMany")
	}

	return nil
}
