package routes

import (
	"my-store/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Echo) {
	e.POST("/product", controllers.CreateProduct)
	// e.GET("/user/:userId", controllers.GetUserById)
	// e.PUT("/user/:userId", controllers.EditAUser)
	// e.DELETE("/user/:userId", controllers.DeleteAUser)
	// e.GET("/users", controllers.GetAllUsers)
}
