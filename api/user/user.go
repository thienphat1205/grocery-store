package user

import (
	"context"
	"my-store/pkg/models"
)

type UserService interface {
	CreateUser(ctx context.Context, request *CreateUserReq) (*CreateUserResp, error)
	GetUserById(ctx context.Context, request *GetUserByIdReq) (*models.User, error)
}

type GetUserByIdReq struct {
	Id string `json:"id,omitempty"  validate:"required"`
}

type CreateUserReq struct {
	Name     string `json:"name,omitempty"  validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
	Store    string `json:"storeId,omitempty"`
}

type CreateUserResp struct {
}
