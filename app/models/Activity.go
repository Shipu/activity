package models

import (
	"github.com/globalsign/mgo/bson"
	mongo "github.com/go-bongo/bongo"
	_repo "github.com/shipu/tracker/app/repositories"
	"math/big"
)

type Activity struct {
	Event              string   `bson:"event"`
	Model_ID           big.Int  `bson:"model_id"`
	Model_Type         string   `bson:"model_type"`
	UUID               string   `bson:"uuid"`
	Data               []string `json:"data" bson:data`
	Remarks            string   `bson:remarks`
	CreatedId          big.Int  `bson:created_id`
	CreatedType        string   `bson:created_type`
	UpdatedId          big.Int  `bson:updated_id`
	UpdatedType        string   `bson:updated_type`
	mongo.DocumentBase `bson:",inline"`
}

func (activity Activity) GetId() bson.ObjectId {
	return activity.Id
}

// Sets the ID for the document
func (activity Activity) SetId(id bson.ObjectId) {
	activity.Id = id
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
