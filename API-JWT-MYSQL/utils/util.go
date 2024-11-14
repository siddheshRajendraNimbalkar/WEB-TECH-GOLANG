package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CheckBodyPayloag(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing bodys")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func SendErrorPayload(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	return SendErrorPayload(w, status, map[string]string{"error": err.Error()})
}
