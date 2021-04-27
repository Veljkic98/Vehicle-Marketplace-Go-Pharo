package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
)

type CommentController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	// Save(response http.ResponseWriter, request *http.Request)
	// DeleteAll(response http.ResponseWriter, request *http.Request)
}

type commentController struct{}

var (
	commentService service.CommentService
)

func NewCommentController(service service.CommentService) CommentController {
	commentService = service
	return &controller{}
}

func (*commentController) GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	comments, err := commentService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the comments"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comments)
}
