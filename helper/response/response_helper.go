package response

import (
	"encoding/json"
	"net/http"
)

// Error will write error response in JSON format
func Error(w http.ResponseWriter, code int, msg string) {
	JSON(w, code, map[string]string{"message": msg})
}

// JSON will write response in JSON format
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
