package controllers

import (
	"my-store/models"
	"my-store/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func CreateStore(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var store models.Store
	defer cancel()

	//validate the request body
	if err := c.Bind(&store); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StoreResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&store); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.StoreResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newStore := models.Store{
		Id:      primitive.NewObjectID(),
		Name:    store.Name,
		Code:    store.Code,
		Address: store.Address,
	}

	result, err := storeCollection.InsertOne(ctx, newStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.StoreResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetStoreById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	storeId := c.Param("storeId")
	var store models.Store
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(storeId)

	err := storeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&store)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.StoreResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": store}})
}

func EditAStore(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	storeId := c.Param("storeId")
	var store models.Store
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(storeId)
	//validate the request body
	if err := c.Bind(&store); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StoreResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&store); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.StoreResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := models.Store{
		Name:    store.Name,
		Code:    store.Code,
		Address: store.Address,
	}

	result, err := storeCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//get updated store details
	var updatedStore models.Store
	if result.MatchedCount == 1 {
		err := storeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedStore)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.StoreResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updatedStore}})
}

func DeleteAStore(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	storeId := c.Param("storeId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(storeId)

	result, err := storeCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.StoreResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Store with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.StoreResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Store successfully deleted!"}})
}

func GetAllStores(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var stores []models.Store
	defer cancel()

	results, err := storeCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var store models.Store
		if err = results.Decode(&store); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.StoreResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		stores = append(stores, store)
	}

	return c.JSON(http.StatusOK, responses.StoreResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": stores}})
}