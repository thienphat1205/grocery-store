package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" validate:"required"`
	Code        string             `bson:"code,omitempty" validate:"required"`
	Price       string             `bson:"price,omitempty" validate:"required"`
	Description string             `bson:"description,omitempty"`
	CategoryId  string             `bson:"categoryId,omitempty"`
}
