package api

import (
	_model "github.com/shipu/tracker/app/models"
	_repo "github.com/shipu/tracker/app/repositories"
	"github.com/shipu/tracker/app/response"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func Read(w http.ResponseWriter, r *http.Request)  {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func All(w http.ResponseWriter, r *http.Request)  {
	repo := _repo.Create()

	response.WithJson(w, repo, 200)
}

func LastActivity(w http.ResponseWriter, r *http.Request)  {

	ctx := r.Context()
	//activity := ctx.Value("activity")
	activity, ok := ctx.Value("activity").(*_model.Activity)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	response.WithJson(w, activity, 200)
}


