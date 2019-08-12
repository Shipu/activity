package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Activity struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:_id,omitempty`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}