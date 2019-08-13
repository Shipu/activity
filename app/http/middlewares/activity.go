package middlewares

import (
	"context"
	"github.com/go-chi/chi"
	_model "github.com/shipu/tracker/app/models"
	_repo "github.com/shipu/tracker/app/repositories"
	"github.com/shipu/tracker/app/response"
	"net/http"
)

func ActivityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activityId := chi.URLParam(r, "activityId")

		activity, err := _repo.Model(_model.Activity{}).FindById(activityId)

		if err != nil {
			response.WithError(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "activity", activity)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

