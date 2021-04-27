package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
)

type OfferController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	// Save(response http.ResponseWriter, request *http.Request)
	// DeleteAll(response http.ResponseWriter, request *http.Request)
}

type offerController struct{}

var (
	offerService service.OfferService
)

func NewOfferController(service service.OfferService) OfferController {
	offerService = service
	return &offerController{}
}

func (*offerController) GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	offers, err := offerService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the offers."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(offers)
}
