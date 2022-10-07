package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	"my-store/internal/log"
	"my-store/pkg/routes"
	"my-store/pkg/server"
)

func main() {

	e := echo.New()

	// init logger
	if err := log.Init(); err != nil {
		panic(err)
	}

	//run database

	gServer, _ := server.NewServiceContext()

	// if err != nil {
	// 	// fmt.Println("initiate server context failed", zap.Error(err))
	// }
	fmt.Println(gServer)
	//routes
	routes.UserRoute(e.Group("/user"))
	routes.ProductRoute(e.Group("/product"))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
