package product

import (
	"context"
)

type ProductService interface {
	GetProductById(ctx context.Context, request *GetProductByIdReq) (*GetProductByIdResp, error)
}

type GetProductByIdReq struct {
	ProductId string `json:"productId"`
}

type GetProductByIdResp struct {
	Data *ProductDetail `json:"data"`
}

type ProductDetail struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}
