package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.uber.org/zap"
)

type Database interface {
	Collection(ctx context.Context, name string) *mongo.Collection
}

type MongoEntity interface {
}

type Instance interface {
	SessionManager
	Client() *mongo.Client
	Close()
	Database(name string, opts ...*options.DatabaseOptions) Database
	Collection(ctx context.Context, dbName, collName string) *mongo.Collection
}

type SessionManager interface {
	WithSession(ctx context.Context, fn func(context.Context) error) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

func newInstance(client *mongo.Client, useTransaction bool) Instance {
	return &mongoClient{
		client:         client,
		useTransaction: useTransaction,
	}
}

type mongoClient struct {
	logger         *zap.Logger
	client         *mongo.Client
	useTransaction bool
}

func (rcv *mongoClient) Close() {
	rcv.client.Disconnect(context.Background())
}

func (rcv *mongoClient) Database(name string, _ ...*options.DatabaseOptions) Database {
	return &database{
		client: rcv,
		name:   name,
	}
}

func (rcv *mongoClient) Collection(ctx context.Context, dbName, collName string) *mongo.Collection {
	if sess := mongo.SessionFromContext(ctx); sess != nil {
		return sess.Client().Database(dbName).Collection(collName)
	}
	return rcv.Client().Database(dbName).Collection(collName)
}

func (rcv *mongoClient) Client() *mongo.Client {
	return rcv.client
}

type contextTransaction struct {
}

func (rcv *mongoClient) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	if !rcv.useTransaction {
		return fn(ctx)
	}
	if v := ctx.Value(contextTransaction{}); v != nil {
		return fn(ctx)
	}

	opts := options.Transaction().
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	// if mongodb's session context exists
	if sess := mongo.SessionFromContext(ctx); sess != nil {
		_, err := sess.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
			return nil, fn(sessCtx)
		}, opts)
		return err
	}

	ctx = context.WithValue(ctx, contextTransaction{}, true)

	return rcv.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		_, err := sessionContext.WithTransaction(sessionContext, func(sessCtx mongo.SessionContext) (interface{}, error) {

			return nil, fn(sessCtx)
		}, opts)
		return err
	})
}

func (rcv *mongoClient) WithSession(ctx context.Context, fn func(context.Context) error) error {
	// if mongodb's session context is already existed
	if sc, ok := ctx.(mongo.SessionContext); ok {
		return fn(sc)
	}

	//opts := options.Session()
	//SetCausalConsistency(true)
	return rcv.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		return fn(sessionContext)
	})
}

type database struct {
	client Instance
	name   string
}

func (db *database) Collection(ctx context.Context, name string) *mongo.Collection {
	return db.client.Collection(ctx, db.name, name)
}
