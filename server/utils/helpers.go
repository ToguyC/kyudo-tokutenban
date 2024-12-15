package utils

import (
	"encoding/json"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}
