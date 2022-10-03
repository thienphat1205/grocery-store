package controllers

import (
	"my-store/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var storeCollection *mongo.Collection = configs.GetCollection(configs.DB, "stores")
var validate = validator.New()
