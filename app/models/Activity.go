package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	mongo "github.com/go-bongo/bongo"
	_repo "github.com/shipu/tracker/app/repositories"
	"math/big"
)

type Activity struct {
	Event              string   `bson:"event"`
	ModelID            big.Int  `bson:"model_id"`
	ModelType          string   `bson:"model_type"`
	UUID               string   `bson:"uuid"`
	Data               []string `json:"data" bson:data`
	Remarks            string   `bson:remarks`
	Url                string   `bson:url`
	UserAgent          []string `json:"user_agent" bson:"user_agent"`
	CreatedId          big.Int  `bson:created_id`
	CreatedType        string   `bson:created_type`
	UpdatedId          big.Int  `bson:updated_id`
	UpdatedType        string   `bson:updated_type`
	mongo.DocumentBase `bson:",inline"`
}

func (activity Activity) BeforeSave(collection *mongo.Collection) error {
	fmt.Println("beforeSave")

	err := *new(error)

	return err
}

func (activity Activity) Validate(collection *mongo.Collection) []error {
	var err []error

	fmt.Println("validation")

	//err = append(err, errors.New("event: Event Required"))

	return  err
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
