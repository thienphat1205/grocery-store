package main

import (
	"my-store/configs"
	"os"

	"github.com/labstack/echo/v4"

	"my-store/routes"
)

func main() {

	e := echo.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e)
	routes.StoreRoute(e)
	routes.ProductRoute(e)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
