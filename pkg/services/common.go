package services

import (
	"my-store/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var Validate = validator.New()

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var ProductCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
