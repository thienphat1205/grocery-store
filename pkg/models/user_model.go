package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"user_name"`
	Location string             `bson:"user_location"`
	Title    string             `bson:"user_title"`
	Store    string             `bson:"store_id"`
}
