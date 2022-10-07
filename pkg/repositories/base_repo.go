package repositories

import (
	"context"
	"sync"
	"time"

	"my-store/pkg/utils"

	"my-store/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepo interface {
	BaseIncrease(ctx context.Context, filter interface{}, fieldName string, value int) error
	Collection(ctx context.Context) *mongo.Collection
	IsNotFound(error) bool
	IsDuplicateKeyError(err error) bool
	BaseCount(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	BaseUpdateOne(ctx context.Context, filter interface{}, data interface{}) error
	BaseUpdateManyObject(ctx context.Context, filter interface{}, data interface{}) error
	BaseUpdateManyData(ctx context.Context, filter interface{}, data bson.M) (int64, error)
	BaseDecodeFirst(ctx context.Context, cursor *mongo.Cursor) (map[string]interface{}, error)
	BaseInsertMany(ctx context.Context, document []interface{}) error
}

type BaseRepoImpl struct {
	db            database.Database
	name          string
	lock          *sync.Mutex
	counter       int64
	lastCountTime int64
}

func (repo *BaseRepoImpl) Init(db database.Database, name string) *BaseRepoImpl {
	repo.db = db
	repo.name = name
	repo.lock = &sync.Mutex{}
	return repo
}

func (repo BaseRepoImpl) Collection(ctx context.Context) *mongo.Collection {
	return repo.db.Collection(ctx, repo.name)
}

func (repo BaseRepoImpl) CollectionConsumed(ctx context.Context) *mongo.Collection {
	return repo.db.Collection(ctx, repo.name+"_consumed")
}

func (repo BaseRepoImpl) IsNotFound(err error) bool {
	return false
	//return errcode.IsNotFoundError(err)
}

func (repo BaseRepoImpl) IsDuplicateKeyError(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

// BaseFind Find list records of entity based on filter and sort criteria input, @parameter sort must not contain reverse filter by `_id` on it.
// And it just should contain sorting on some other columns, @param reverse will reverse the results on column `_id`.
func (repo BaseRepoImpl) BaseFind(ctx context.Context, filter interface{}, offset, limit int, reverse bool, sort bson.D) (*mongo.Cursor, error) {
	findOptions := options.Find()
	if limit == 0 {
		limit = 20
	}
	if limit > 1000 {
		limit = 1000
	}
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	sortCriteria := buildSortCriteria(sort, reverse)
	if len(sortCriteria) > 0 {
		findOptions.SetSort(sortCriteria)
	}
	return repo.Collection(ctx).Find(ctx, filter, findOptions)
}

func (repo BaseRepoImpl) BaseFindAdvance(ctx context.Context, filter interface{}, findOptions ...*options.FindOptions) (*mongo.Cursor, error) {
	return repo.Collection(ctx).Find(ctx, filter, findOptions...)
}

// BaseCount return the number of documents in the collection based on filter.
// The limit of this method will limit the maximum number of documents to count, leave default is 0 here for no limit.
// The offset of this method will skip number of documents before counting. Also let it default for start from beginning.
func (repo BaseRepoImpl) BaseCount(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	finalOption := mergeCountOptions(opts)

	// Can add hint for index column and limit execution time here when needed
	return repo.Collection(ctx).CountDocuments(ctx, filter, finalOption)
}

func (repo BaseRepoImpl) BaseFindNoLimit(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	cursor, err := repo.Collection(ctx).Find(ctx, filter, opts...)
	if err != nil {
		//log.Logger(ctx).Error("fail to find data", zap.String("collection_name", repo.name), zap.Error(err))
		return nil, err
	}
	return cursor, nil
}

func (repo BaseRepoImpl) BaseFindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	return repo.Collection(ctx).FindOne(ctx, filter)
}

func (repo BaseRepoImpl) BaseInsert(ctx context.Context, data interface{}) error {
	_, err := repo.Collection(ctx).InsertOne(ctx, data)
	if err != nil {
		//log.Logger(ctx).Error("fail to insert", zap.String("collection_name", repo.name), zap.Error(err))
		return err
	}
	return nil
}

func (repo BaseRepoImpl) BaseIncrease(ctx context.Context, filter interface{}, fieldName string, value int) error {
	updater := bson.M{
		fieldName: value,
	}

	optionsUpdate := options.Update()
	optionsUpdate.SetUpsert(true)
	_, err := repo.Collection(ctx).UpdateOne(ctx, filter, bson.M{"$inc": updater}, optionsUpdate)
	return err
}

func (repo BaseRepoImpl) BaseUpdateOne(ctx context.Context, filter interface{}, data interface{}) error {
	bUpdater, err := utils.ConvertToBson(data)
	if err != nil {
		return err
	}
	bUpdater["last_updated_time"] = time.Now()

	result, err := repo.Collection(ctx).UpdateOne(ctx, filter, bson.M{"$set": bUpdater})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		//return errcode.Error(errcode.NotFound)
	}

	return nil
}

func (repo BaseRepoImpl) BaseUpdateManyObject(ctx context.Context, filter interface{}, data interface{}) error {
	bUpdater, err := utils.ConvertToBson(data)
	if err != nil {
		return err
	}
	bUpdater["last_updated_time"] = time.Now()

	rs, err := repo.Collection(ctx).UpdateMany(ctx, filter, bson.M{"$set": bUpdater})
	if err != nil {
		return err
	}

	if rs.MatchedCount == 0 {
		//return errcode.Error(errcode.NotFound)
	}

	return nil
}

func (repo BaseRepoImpl) BaseUpdateManyData(ctx context.Context, filter interface{}, data bson.M) (int64, error) {
	data["last_updated_time"] = time.Now()

	rs, err := repo.Collection(ctx).UpdateMany(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return 0, err
	}
	return rs.ModifiedCount, nil
}

func (repo BaseRepoImpl) BaseReplace(ctx context.Context, filter interface{}, data interface{}) error {
	bUpdater, err := utils.ConvertToBson(data)
	if err != nil {
		return err
	}
	bUpdater["last_updated_time"] = time.Now()

	result, err := repo.Collection(ctx).ReplaceOne(ctx, filter, data)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		//return errcode.Error(errcode.NotFound)
	}

	return nil
}

func (repo BaseRepoImpl) BaseFindOneAndUpdate(ctx context.Context, filter interface{}, data interface{}) *mongo.SingleResult {
	bUpdater, err := utils.ConvertToBson(data)
	if err != nil {
		return nil
	}
	bUpdater["last_updated_time"] = time.Now()

	return repo.Collection(ctx).FindOneAndUpdate(ctx, filter, bson.M{"$set": bUpdater})
}

// BaseUpsertOne executes an update command to update at most one document in the collection.
// If the filter not matched with any documents, a new record will be inserted
func (repo BaseRepoImpl) BaseUpsertOne(ctx context.Context, filter interface{}, data interface{}) error {
	optionsUpdate := options.Update()
	optionsUpdate.SetUpsert(true)
	bUpdater, err := utils.ConvertToBson(data)
	if err != nil {
		return err
	}
	bUpdater["last_updated_time"] = time.Now()
	_, err = repo.Collection(ctx).UpdateOne(ctx, filter, bson.M{"$set": bUpdater}, optionsUpdate)
	return err
}

func (repo BaseRepoImpl) BaseDelete(ctx context.Context, filter interface{}) error {
	_, err := repo.Collection(ctx).DeleteOne(ctx, filter)
	if err != nil {
		//log.Logger(ctx).Error("cant delete", zap.Error(err))
		return err
	}
	return nil
}

func (repo BaseRepoImpl) BaseDeleteMany(ctx context.Context, filter interface{}) error {
	_, err := repo.Collection(ctx).DeleteMany(ctx, filter)
	return err
}

func buildSortCriteria(sort bson.D, reverse bool) bson.D {
	var sortCriteria bson.D
	if sort != nil {
		sortCriteria = sort
	} else {
		sortCriteria = bson.D{}
	}
	if reverse {
		sortCriteria = append(sortCriteria, bson.E{Key: "_id", Value: -1})
	}
	return sortCriteria
}

func (repo BaseRepoImpl) GetCounter(time int64) int64 {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	if time > repo.lastCountTime {
		repo.lastCountTime = time
		repo.counter = 1
	} else {
		repo.counter++
	}

	return time + repo.counter
}

func (repo BaseRepoImpl) BaseAggregate(ctx context.Context, query interface{}) (*mongo.Cursor, error) {
	return repo.Collection(ctx).Aggregate(ctx, query)
}

func mergeCountOptions(opts []*options.CountOptions) *options.CountOptions {
	var finalOption = options.Count()

	for _, option := range opts {
		if option.Limit != nil && *option.Limit > 0 {
			finalOption.Limit = option.Limit
		}

		if option.Skip != nil && *option.Skip > 0 {
			finalOption.Skip = option.Skip
		}

		if option.Collation != nil {
			finalOption.Collation = option.Collation
		}

		if option.MaxTime != nil {
			finalOption.MaxTime = option.MaxTime
		}
	}

	return finalOption
}

func (repo BaseRepoImpl) BaseDecodeFirst(ctx context.Context, cursor *mongo.Cursor) (map[string]interface{}, error) {
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	if cursor.Next(ctx) {
		data := make(map[string]interface{})
		err := cursor.Decode(data)
		if err != nil {
			//log.Logger(ctx).Error("fail to decode", zap.String("collection_name", repo.name), zap.Error(err))
			return nil, err
		}
		return data, nil
	}
	return nil, mongo.ErrNoDocuments
}

func (repo BaseRepoImpl) BaseInsertMany(ctx context.Context, documents []interface{}) error {
	_, err := repo.Collection(ctx).InsertMany(ctx, documents)
	if err != nil {
		//log.Logger(ctx).Error("fail to insert many", zap.String("collection_name", repo.name), zap.Error(err))
	}
	return err
}
