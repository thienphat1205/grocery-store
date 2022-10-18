package repositories

import (
	"context"

	"my-store/internal/database"
	"my-store/internal/log"
	"my-store/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type UserRepo interface {
	BaseRepo
	Create(ctx context.Context, ticket *models.User) error
	Get(ctx context.Context, ticketId string) (*models.User, error)
}

func NewUserRepo(dbIns database.Instance, dbName, collName string) UserRepo {
	return &userRepo{
		baseRepo{
			dbIns:    dbIns,
			dbName:   dbName,
			collName: collName,
		},
	}
}

type userRepo struct {
	baseRepo
}

func (rcv userRepo) Create(ctx context.Context, user *models.User) error {
	_, err := rcv.collection(ctx).InsertOne(ctx, user)
	if err != nil {
		log.Logger(ctx).Error("encounter error while inserting user", zap.Error(err))
		return err
	}
	return nil
}

func (rcv userRepo) Get(ctx context.Context, userId string) (*models.User, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{
		"_id": objId,
	}
	var user = new(models.User)
	err := rcv.collection(ctx).FindOne(ctx, filter).Decode(user)
	if err != nil {
		if !rcv.IsNotFound(err) {
			log.Logger(ctx).Error("fail to find user by user code", zap.Error(err))
		}
		return nil, err
	}
	return user, nil
}
