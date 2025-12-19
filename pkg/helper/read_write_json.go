package helper

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"social/pkg/env"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

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

	err := dec.Decode(data)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("request body must not be empty")
		}
		return err
	}

	return nil
}

func WriteErrorJson(w http.ResponseWriter, status int, err error) error {
	logError(err)

	var ce *CustomError
	if errors.As(err, &ce) {
		return WriteJson(w, status, ErrorResponse{ErrorBody{Message: ce.message, Details: ce.details}})
	}

	return WriteJson(w, status, ErrorResponse{
		ErrorBody{Message: err.Error(), Details: nil},
	})
}

func logError(err error) {
	var ce *CustomError
	if errors.As(err, &ce) {
		if ce.cause != nil {
			log.Printf(
				"[ERROR] msg=%s cause=%v",
				ce.message,
				ce.cause,
			)
			return
		}

		log.Printf("[ERROR] msg=%s", ce.message)
		return
	}

	log.Printf("[ERROR] %v", err)
}
