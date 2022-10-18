package user

import (
	"context"
	api "my-store/api/user"
	"my-store/pkg/models"
	"my-store/pkg/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserService(factory repositories.Factory) api.UserService {
	return &sortingIssueService{
		userRepo: factory.UserRepo(),
	}
}

type sortingIssueService struct {
	userRepo repositories.UserRepo
}

func (rcv sortingIssueService) CreateUser(ctx context.Context, request *api.CreateUserReq) (*api.CreateUserResp, error) {

	//validate the request body

	newUser := &models.User{
		Id:       primitive.NewObjectID(),
		Name:     request.Name,
		Location: request.Location,
		Title:    request.Title,
		Store:    request.Store,
	}

	err := rcv.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (rcv sortingIssueService) GetUserById(ctx context.Context, request *api.GetUserByIdReq) (*models.User, error) {
	user, err := rcv.userRepo.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
