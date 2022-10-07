package user

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, request *CreateUserReq) (*CreateUserResp, error)
}

type CreateUserReq struct {
	Name     string `json:"name,omitempty"  validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
	Store    string `json:"storeId,omitempty"`
}

type CreateUserResp struct {
}
