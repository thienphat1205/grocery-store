package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"  bson:"store_name,omitempty" validate:"required"`
	Code    string             `json:"code,omitempty" bson:"store_code,omitempty" validate:"required"`
	Address string             `json:"address,omitempty" bson:"store_address,omitempty" validate:"required"`
}
