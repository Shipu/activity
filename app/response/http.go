package response

import (
	"encoding/json"
	"net/http"
)

func WithJson(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func WithError(w http.ResponseWriter, message string, code int) {
	WithJson(w, map[string]string{"message": message}, code)
}