package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson[T any](r *http.Request, payload T) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, err)
}
