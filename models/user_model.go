package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"  bson:"user_name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" bson:"user_location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" bson:"user_title,omitempty" validate:"required"`
	Store    string             `json:"storeId,omitempty" bson:"store_id,omitempty" validate:"required"`
}
