package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func Format(err error) map[string]string {
	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errors[field] = message(err)
	}

	return errors
}

func message(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "min":
		return "is too short"
	case "max":
		return "is too long"
	case "gt":
		return "must be greater than 0"
	default:
		return "is invalid"
	}
}
