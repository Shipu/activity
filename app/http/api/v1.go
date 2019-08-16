package api

import (
	"encoding/json"
	mongo "github.com/go-bongo/bongo"
	_model "github.com/shipu/tracker/app/models"
	_repo "github.com/shipu/tracker/app/repositories"
	"github.com/shipu/tracker/app/response"
	"io/ioutil"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var activity _model.Activity
	err = json.Unmarshal(body, &activity)

	if err != nil {
		response.WithError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = _model.ActivityRepo().Save(activity)

	if err != nil {

		if vErr, ok := err.(*mongo.ValidationError); ok {
			response.WithMultipleError(w, vErr.Errors, http.StatusBadRequest)
			log.Println("Validation errors are:", vErr.Errors)
		} else {
			log.Println("Got a real error:", err.Error())
		}

		return
	}

	response.WithJson(w, activity, http.StatusCreated)

}

func Read(w http.ResponseWriter, r *http.Request) {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func Update(w http.ResponseWriter, r *http.Request) {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func All(w http.ResponseWriter, r *http.Request) {
	repo := _model.ActivityRepo().All()

	response.WithJson(w, repo, 200)
}

func LastActivity(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	activity, ok := ctx.Value("activity").(*_model.Activity)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	response.WithJson(w, activity, 200)
}
