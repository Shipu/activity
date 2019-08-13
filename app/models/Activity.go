package models

import (
	"github.com/go-bongo/bongo"
)

type Activity struct {
	Name string `bson:"name"`
	Email string `bson:"email"`
	bongo.DocumentBase `bson:",inline"`
}

func (activity Activity) Collection() string {
	return "activity_logs"
}

func (activity Activity) Instance() interface{} {
	return &Activity{}
}


