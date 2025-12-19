package validator

import (
	"net/http"
	"social/pkg/helper"
)

func PayloadValidator[T any](w http.ResponseWriter, req *http.Request, payload *T) bool {
	if err := helper.ReadJSON(w, req, &payload); err != nil {
		helper.WriteErrorJson(w, http.StatusBadRequest, err)
		return false
	}

	if err := Validate.Struct(payload); err != nil {
		helper.WriteErrorJson(
			w,
			http.StatusBadRequest,
			helper.NewCustomError(err,
				"invalid input",
				Format(err)),
		)
		return false
	}

	return true
}
