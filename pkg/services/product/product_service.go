package product

import (
	"context"
	"fmt"
	api "my-store/api/product"
	"my-store/pkg/repositories"
)

func ProductService(factory repositories.Factory) api.ProductService {
	return &sortingIssueService{
		number: 10,
	}
}

type sortingIssueService struct {
	number int
}

func (rcv sortingIssueService) GetProductById(ctx context.Context, request *api.GetProductByIdReq) (*api.GetProductByIdResp, error) {
	productId := request.ProductId
	fmt.Println(productId)
	// productId := request.ProductId
	// var data *api.ProductDetail
	// objId, _ := primitive.ObjectIDFromHex(productId)
	// err := services.ProductCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&data)
	// if err != nil {
	// 	return nil, err
	// }
	return &api.GetProductByIdResp{
		Data: nil,
	}, nil
}
