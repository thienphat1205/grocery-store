package services

import (
	"my-store/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
