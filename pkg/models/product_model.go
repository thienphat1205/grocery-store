package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"  bson:"name,omitempty" validate:"required"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty" validate:"required"`
	Price       string             `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CategoryId  string             `json:"categoryId,omitempty" bson:"categoryId,omitempty"`
}
