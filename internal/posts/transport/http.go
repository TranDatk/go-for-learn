package transport

import (
	"net/http"
	"social/internal/entity"
	"social/internal/posts/service"
	"social/pkg/helper"
	"social/pkg/validator"
)

type PostHandler struct {
	svc service.IPostService
}

func NewPostHandler(svc service.IPostService) *PostHandler {
	return &PostHandler{svc: svc}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, req *http.Request) {
	param := req.PathValue("id")

	ctx := req.Context()

	post, err := h.svc.GetPostByID(ctx, param)

	if err != nil {
		helper.WriteErrorJson(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJson(w, http.StatusOK, post)
}

func (h *PostHandler) NewPost(w http.ResponseWriter, req *http.Request) {
	var payload entity.CreatePostPayload

	if !validator.PayloadValidator(w, req, &payload) {
		return
	}

	post := entity.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
	}

	ctx := req.Context()

	if err := h.svc.CreatePost(ctx, &post); err != nil {
		helper.WriteErrorJson(w, http.StatusInternalServerError, err)
		return
	}
}
