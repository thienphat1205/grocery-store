package repositories

import (
	"context"
	"fmt"
	"my-store/internal/database"
	"my-store/internal/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type BaseRepo interface {
	collection(ctx context.Context) *mongo.Collection
	IsNotFound(err error) bool
}

type baseRepo struct {
	dbIns    database.Instance
	dbName   string
	collName string
}

func (rcv baseRepo) batchSize(limit int64) *options.FindOptions {
	if limit > 1000 {
		limit = 1000
	}
	return options.Find().SetBatchSize(int32(limit))
}

func (rcv baseRepo) collection(ctx context.Context) *mongo.Collection {
	return rcv.dbIns.Collection(ctx, rcv.dbName, rcv.collName)
}

func (rcv baseRepo) IsNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}

func (rcv baseRepo) mapInCursor(ctx context.Context, cursor *mongo.Cursor) ([]map[string]interface{}, error) {
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var data = make([]map[string]interface{}, 0)
	if cursor.Next(ctx) {
		var item = make(map[string]interface{})
		if err := cursor.Decode(item); err != nil {
			log.Logger(ctx).Error("fail to decode data", zap.Error(err))
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (rcv baseRepo) firstInCursor(ctx context.Context, cursor *mongo.Cursor, data map[string]interface{}) error {
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	if cursor.Next(ctx) {
		return cursor.Decode(data)
	}
	return nil
}

func (rcv baseRepo) sumField(ctx context.Context, pipeline []bson.M, field string) (float64, error) {
	pipeline = append(pipeline, bson.M{
		"$group": bson.M{
			"_id": "total",
			"total": bson.M{
				"$sum": fmt.Sprintf("$%s", field),
			},
		},
	})
	cursor, err := rcv.collection(ctx).Aggregate(ctx, pipeline)
	if err != nil {
		log.Logger(ctx).Error("fail to sum field: " + err.Error())
		return 0, err
	}

	var data = make(map[string]interface{})
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	if cursor.Next(ctx) {
		err = cursor.Decode(data)
		if err != nil {
			return 0, err
		}
	}
	switch data["total"].(type) {
	case int64:
		return float64(data["total"].(int64)), nil
	case float64:
		return data["total"].(float64), nil
	case float32:
		return float64(data["total"].(float32)), nil
	case int32:
		return float64(data["total"].(int32)), nil
	case int:
		return float64(data["total"].(int)), nil
	}
	return 0, nil
}
