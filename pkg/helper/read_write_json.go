package helper

import (
	"encoding/json"
	"net/http"
	"social/pkg/env"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := env.Get("MAX_BYTES", 1048576)

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(data)
}

func WriteErrorJson(w http.ResponseWriter, status int, message any) error {
	return WriteJson(w, status, message)
}
