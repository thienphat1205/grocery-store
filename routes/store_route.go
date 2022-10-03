package routes

import (
	"my-store/controllers"

	"github.com/labstack/echo/v4"
)

func StoreRoute(e *echo.Echo) {
	e.POST("/store", controllers.CreateStore)
	e.GET("/store/:storeId", controllers.GetStoreById)
	e.PUT("/store/:storeId", controllers.EditAStore)
	e.DELETE("/store/:storeId", controllers.DeleteAStore)
	e.GET("/stores", controllers.GetAllStores)
}
