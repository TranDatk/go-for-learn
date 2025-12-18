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

func (h *PostHandler) NewPostService(w http.ResponseWriter, req *http.Request) {
	var payload entity.CreatePostPayload

	if err := helper.ReadJSON(w, req, &payload); err != nil {
		helper.WriteErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(payload); err != nil {
		helper.WriteErrorJson(
			w,
			http.StatusBadRequest,
			validator.Format(err),
		)
		return
	}

	post := entity.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserID,
		Tags:    payload.Tags,
	}

	ctx := req.Context()

	if err := h.svc.CreatePost(ctx, &post); err != nil {
		helper.WriteErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

}
