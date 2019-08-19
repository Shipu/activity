package response

import (
	"encoding/json"
	"net/http"
)

func WithJson(w http.ResponseWriter, payload interface{}, code int) {
	jsonResponse(w, payload, code, false)
}

func WithError(w http.ResponseWriter, message string, code int) {
	jsonResponse(w, map[string]string{"message": message}, code, true)
}

func WithMultipleError(w http.ResponseWriter, payload interface{}, code int) {
	jsonResponse(w, payload, code, true)
}


func jsonResponse(w http.ResponseWriter, payload interface{}, code int, error bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if !error {
		payload, _ = json.Marshal(payload)
	}

	apiResponse := map[string]interface{}{
		"code": code,
		"data": payload,
	}

	response, _ := json.Marshal(apiResponse)

	w.Write(response)
}
