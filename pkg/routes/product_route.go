package routes

import (
	"my-store/pkg/framework/api"
	"my-store/pkg/services/product"

	"github.com/labstack/echo/v4"
)

func ProductRoute(group *echo.Group) {
	productSv := product.ProductService()
	api.AddRoute(group, "/get-by-id", productSv.GetProductById)
}
