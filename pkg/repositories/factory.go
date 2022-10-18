package repositories

import (
	"my-store/internal/database"
)

func NewFactory(mongoClient database.Instance, mainDbName string) Factory {
	return Factory{
		mongoClient: mongoClient,
		mainDbName:  mainDbName,
	}
}

type Factory struct {
	mongoClient database.Instance
	mainDbName  string
}

func (f Factory) Client() database.Instance {
	return f.mongoClient
}

func (f Factory) SessionManager() database.SessionManager {
	return f.mongoClient
}

// ========================== User ==========================

func (f Factory) UserRepo() UserRepo {
	return NewUserRepo(f.mongoClient, f.mainDbName, "users")
}
