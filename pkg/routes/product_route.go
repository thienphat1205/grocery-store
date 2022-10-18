package routes

import (
	"my-store/pkg/framework/api"
	"my-store/pkg/repositories"
	"my-store/pkg/services/product"

	"github.com/labstack/echo/v4"
)

func ProductRoute(group *echo.Group, factory repositories.Factory) {
	productSv := product.ProductService(factory)
	api.AddRoute(group, "/get-by-id", productSv.GetProductById)
}
