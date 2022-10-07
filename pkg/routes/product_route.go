package routes

import (
	"my-store/pkg/framework/api"
	"my-store/pkg/services/product"

	"github.com/labstack/echo/v4"
)

func ProductRoute(publicGroup *echo.Group) {
	productSv := product.SortingIssueService()
	api.AddRoute(publicGroup, "/get-by-id", productSv.GetProductById)
}
