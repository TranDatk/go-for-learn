package helper

import (
	"fmt"
)

type CustomError struct {
	cause   error
	message string
	details any
}

func (e *CustomError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}

func (e *CustomError) Unwrap() error {
	return e.cause
}

func NewCustomError(cause error, msg string, dt any) error {
	return &CustomError{
		cause:   cause,
		message: msg,
		details: dt,
	}
}
