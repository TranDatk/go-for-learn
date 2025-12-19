package transport

import (
	"net/http"
	"social/internal/entity"
	"social/internal/users/service"
	"social/pkg/helper"
	"social/pkg/validator"
)

type UserHandler struct {
	svc service.IUserService
}

func NewUserHandler(svc service.IUserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(w http.ResponseWriter, req *http.Request) {
	var payload entity.CreateUserPayload

	if !validator.PayloadValidator(w, req, &payload) {
		return
	}

	user := &entity.User{
		Name:     payload.Name,
		Username: payload.Username,
		Password: payload.Password,
	}

	ctx := req.Context()

	if err := h.svc.Register(ctx, user); err != nil {
		helper.WriteErrorJson(w, http.StatusInternalServerError, err)
		return
	}

	if err := helper.WriteJson(w, http.StatusCreated, user.ID); err != nil {
		helper.WriteErrorJson(w, http.StatusInternalServerError, err)
		return
	}
}
