package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	mongo "github.com/go-bongo/bongo"
	_i "github.com/shipu/tracker/app/interfaces"
	_db "github.com/shipu/tracker/bootstrap"
)

type Repo struct {
	collection *mongo.Collection
	model      interface{}
}

func Create() string {
	return "from repo create"
}

func (repo Repo) Save(doc mongo.Document) error {
	err := repo.collection.Save(doc)

	return err
}

func (repo Repo) Find(query interface{}) interface{} {
	results := repo.collection.Find(query)

	var items []interface{}

	for results.Next(repo.model) {
		item := repo.model
		items = append(items, item)
	}

	return items
}

func (repo Repo) FindOne(query interface{}) (interface{}, error) {
	result := repo.model
	err := repo.collection.FindOne(query, result)

	return result, err
}

func (repo *Repo) FindById(StringId string) (interface{}, error) {

	result := repo.model
	err := repo.collection.FindById(bson.ObjectIdHex(StringId), result)

	return result, err
}

func (repo *Repo) All() interface{} {
	results := repo.Find(bson.M{})

	return results
}

func (repo Repo) DeleteOne(query bson.M) (error) {
	err := repo.collection.DeleteOne(query)

	return err
}

func (repo Repo) DeleteAll(query bson.M) (*mgo.ChangeInfo, error) {
	changeInfo, err := repo.collection.Delete(query)

	return changeInfo, err
}

func Model(model _i.Model) *Repo {

	collection := _db.Connection.Collection(model.Collection())

	repo := new(Repo)

	repo.collection = collection
	repo.model = model.Instance()

	return repo
}
