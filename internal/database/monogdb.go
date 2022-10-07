package database

import (
	"context"
	"my-store/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongoDb(ctx context.Context) (Instance, error) {
	// logger := log.Logger(ctx)
	// logger.Sugar().Infof("initializing connection to mongodb, username: %s, uri: %s, replSet: %s, hosts: %s, authSource: %s",
	// 	config.Username, strings.TrimSpace(config.URI), config.ReplSet, strings.Join(config.Hosts, ","), config.AuthSource)

	// clientOpts := options.Client().
	// 	SetHosts(config.Hosts).
	// 	SetAuth(options.Credential{
	// 		AuthSource:  config.AuthSource,
	// 		Username:    config.Username,
	// 		Password:    config.Password,
	// 		PasswordSet: true,
	// 	}).
	// 	SetReplicaSet(config.ReplSet)
	// if strings.TrimSpace(config.URI) != "" {
	// 	clientOpts = clientOpts.ApplyURI(config.URI)
	// }

	clientOpts := options.Client().ApplyURI(configs.EnvMongoURI())

	// set timeout to connect mongodb after 10 seconds
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// logger.Info("connecting to mongodb...")

	// connect
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		// logger.Error("connecting mongodb failed", zap.Error(err))
		return nil, err
	}
	// logger.Info("connected to mongodb")

	// logger.Info("ping mongodb...")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		// logger.Error("ping mongodb failed", zap.Error(err))
		return nil, err
	}
	// log.Logger(ctx).Sugar().Infof("ping success mongodb host: %s", clientOpts.Hosts)

	return newInstance(client, true), nil
}

type Config struct {
	URI              string   `json:"uri" yaml:"uri"`
	Hosts            []string `json:"hosts" yaml:"hosts"`
	AuthSource       string   `json:"authSource" yaml:"authSource"`
	Username         string   `json:"username" yaml:"username"`
	Password         string   `json:"password" yaml:"password"`
	ReplSet          string   `json:"replSet" yaml:"replSet"`
	DbName           string   `json:"dbName" yaml:"dbName"`
	UnUseTransaction bool     `json:"unUseTransaction" yaml:"unUseTransaction"`
}
