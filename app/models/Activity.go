package models

import (
	"github.com/go-bongo/bongo"
	_repo "github.com/shipu/tracker/app/repositories"
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

func ActivityRepo() *_repo.Repo {

	return _repo.Model(Activity{})
}


