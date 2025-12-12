package transport

import (
	"net/http"
	"social/internal/posts/service"
)

type PostHandler struct {
	svc service.IPostService
}

func NewPostHandler(svc service.IPostService) *PostHandler {
	return &PostHandler{svc: svc}
}

func (h *PostHandler) NewPostService(w http.ResponseWriter, req *http.Request) {

}
