package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
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
	activity := ctx.Value("activity")
	//if !ok {
	//	http.Error(w, http.StatusText(422), 422)
	//	return
	//}
	response.WithJson(w, fmt.Sprintf("title:%s", activity), 200)
}

func ActivityCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activityId := chi.URLParam(r, "activityId")
		//activity, err := _repo.Find(activityId)
		//err := nil
		//if err != nil {
		//	response.WithError(w, http.StatusText(404), 404)
		//	return
		//}

		//if activityId == string {
		//	response.WithError(w, http.StatusText(404), 404)
		//	return
		//}
		ctx := context.WithValue(r.Context(), "activity", activityId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

