package api

import (
	_repo "github.com/shipu/tracker/app/repositories"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	response := _repo.Create()

	w.Write([]byte(response))
}

func Read(w http.ResponseWriter, r *http.Request)  {
	response := _repo.Create()

	w.Write([]byte(response))
}

func Update(w http.ResponseWriter, r *http.Request)  {

}

func Delete(w http.ResponseWriter, r *http.Request)  {

}

func All(w http.ResponseWriter, r *http.Request)  {
	response := _repo.Create()

	w.Write([]byte(response))
}

func LastActivity(w http.ResponseWriter, r *http.Request)  {
	response := _repo.Create()

	w.Write([]byte(response))
}


