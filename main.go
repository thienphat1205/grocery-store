package main

import (
	"my-store/configs"
	"os"

	"github.com/labstack/echo/v4"

	"my-store/pkg/routes"
)

func main() {

	e := echo.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e.Group("/user"))
	routes.StoreRoute(e)
	routes.ProductRoute(e.Group("/product"))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
