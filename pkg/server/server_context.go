package server

import (
	"context"

	"my-store/internal/database"
	"my-store/internal/log"

	"my-store/pkg/repositories"

	"go.uber.org/zap"
)

// NewServiceContext take environment's configurations to initiate a server context
// which associated database.Database, redis.Client
func NewServiceContext() (server *ServiceContext, err error) {
	ctx := context.Background()
	var mainDb database.Instance

	// logger associated
	logger := log.Global()
	logger.Error("initiating mongodb client")

	// recover if panic
	defer func() {
		if err != nil {
			// logger.Warn("disconnect connections on server context")
			// disconnect database connection
			if mainDb != nil {
				mainDb.Close()
			}
		}
	}()

	// set up databases
	mainDb, err = database.SetupMongoDb(ctx)
	if err != nil {
		return
	}

	// logger.Info("initiating repository factory")
	// set up repository factory
	repositoryFactory := repositories.NewFactory(mainDb, "grocery-store")
	// logger.Info("initiated repository factory")

	return &ServiceContext{
		// logger:  logger,
		mainDb:  mainDb,
		factory: repositoryFactory,
	}, nil
}

type ServiceContext struct {
	logger *zap.Logger
	mainDb database.Instance

	factory repositories.Factory
}

// Shutdown close streams
func (rcv *ServiceContext) Shutdown() {
	if rcv.mainDb != nil {
		rcv.mainDb.Close()
		rcv.logger.Warn("closed mongodb client")
	}

}

func (rcv *ServiceContext) Repositories() repositories.Factory {
	return rcv.factory
}
