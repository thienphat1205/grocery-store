package user

import (
	"context"
	api "my-store/api/user"
)

func UserService() api.UserService {
	return &sortingIssueService{
		number: 10,
	}
}

type sortingIssueService struct {
	number int
}

func (rcv sortingIssueService) CreateUser(ctx context.Context, request *api.CreateUserReq) (*api.CreateUserResp, error) {

	//validate the request body

	// newUser := models.User{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     request.Name,
	// 	Location: request.Location,
	// 	Title:    request.Title,
	// 	Store:    request.Store,
	// }

	// _, err := services.UserCollection.InsertOne(ctx, newUser)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
