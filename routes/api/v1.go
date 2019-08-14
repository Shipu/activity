package api

import (
	"github.com/go-chi/chi"
	_http "github.com/shipu/tracker/app/http/api"
	_middleware "github.com/shipu/tracker/app/http/middlewares"
	"net/http"
)

func Load() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Activity Crud using Golang, Mongodb, Chi, Bongo, GodotEnv"))
	})

	r.Route("/v1", func(r chi.Router) {

		r.Get("/activities", _http.All)

		r.Post("/activity", _http.Create)

		r.Route("/{activityId}", func(r chi.Router) {
			r.Use(_middleware.ActivityMiddleware)
			r.Get("/", _http.Read)
			r.Put("/", _http.Update)
			r.Delete("/", _http.Delete)

			r.Get("/last/activity", _http.LastActivity)
		})

	})

	return r
}
