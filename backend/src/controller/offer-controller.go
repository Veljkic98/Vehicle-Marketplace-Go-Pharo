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

func NewOfferController(service service.VehicleService) OfferController {
	vehicleService = service
	return &controller{}
}

func (*offerController) GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	vehicles, err := vehicleService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the vehicles"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(vehicles)
}
