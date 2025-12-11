package helper

import "fmt"

type RequestError struct {
	Msg   string
	Cause error
}

func (e *RequestError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Msg, e.Cause)
	}
	return e.Msg
}

func (e *RequestError) Unwrap() error {
	return e.Cause
}

func NewRequestError(msg string, cause error) error {
	return &RequestError{
		Msg:   msg,
		Cause: cause,
	}
}
