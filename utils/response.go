package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseError return error message
func ResponseError(w http.ResponseWriter, code int, msg string) {
	ResponseJSON(w, code, map[string]string{"message": msg})
}

// ResponseJSON write JSON response format
func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
