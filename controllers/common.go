package controllers

import (
	"my-store/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var storeCollection *mongo.Collection = configs.GetCollection(configs.DB, "stores")
var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var inventoryManagementCollection *mongo.Collection = configs.GetCollection(configs.DB, "inventory_management")
var importWarehouseCollection *mongo.Collection = configs.GetCollection(configs.DB, "import_warehouse")
var exportWarehouseCollection *mongo.Collection = configs.GetCollection(configs.DB, "export_warehouse")
var validate = validator.New()
