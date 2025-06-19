package response

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var webResponse interface{}

	switch status {
	case 200:
		webResponse = struct {
			StatusCode int    `json:"StatusCode"`
			Message    string `json:"message,omitempty"`
			Data       any    `json:"data,omitempty"`
		}{
			StatusCode: status,
			Message:    message,
			Data:       data,
		}
	case 404, 422, 500:
		webResponse = struct {
			StatusCode int    `json:"StatusCode"`
			Message    string `json:"message,omitempty"`
		}{
			StatusCode: status,
			Message:    message,
		}
	}

	json.NewEncoder(w).Encode(&webResponse)
}
