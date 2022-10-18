package routes

import (
	"my-store/pkg/framework/api"
	"my-store/pkg/repositories"
	"my-store/pkg/services/user"

	"github.com/labstack/echo/v4"
)

func UserRoute(group *echo.Group, factory repositories.Factory) {
	userSv := user.UserService(factory)
	api.AddRoute(group, "/create", userSv.CreateUser)
	api.AddRoute(group, "/get-by-id", userSv.GetUserById)
}

// func UserRoute(e *echo.Echo) {
// 	e.POST("/user", controllers.CreateUser)
// 	e.GET("/user/:userId", controllers.GetUserById)
// 	e.PUT("/user/:userId", controllers.EditAUser)
// 	e.DELETE("/user/:userId", controllers.DeleteAUser)
// 	e.GET("/users", controllers.GetAllUsers)
// }
