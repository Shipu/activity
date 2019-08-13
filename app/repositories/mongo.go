package repositories

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	mongo "github.com/go-bongo/bongo"
	_i "github.com/shipu/tracker/app/interfaces"
	_db "github.com/shipu/tracker/bootstrap"
)


type Repo struct {
	collection *mongo.Collection
	model interface{}
}


func Create() string {
	return "from repo create"
}


func (repo *Repo) FindById(StringId string) (interface{}, error){

	result := repo.model
	err := repo.collection.FindById(bson.ObjectIdHex(StringId), result)

	return result, err
}

func (repo *Repo) All() interface{} {
	results := repo.collection.Find(bson.M{})

	var items []interface{}

	for results.Next(repo.model) {
		items = append(items, repo.model)
		fmt.Println(repo.model)
	}

	return items
}


func Model(model _i.Model) *Repo {

	collection := _db.Connection.Collection(model.Collection())

	repo := new(Repo)

	repo.collection = collection
	repo.model = model.Instance()

	return repo
}

