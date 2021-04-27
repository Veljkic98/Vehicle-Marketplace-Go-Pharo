package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
)

type RateController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	// Save(response http.ResponseWriter, request *http.Request)
	// DeleteAll(response http.ResponseWriter, request *http.Request)
}

type rateController struct{}

var (
	rateService service.RateService
)

func NewRateController(service service.RateService) RateController {
	rateService = service
	return &controller{}
}

func (*rateController) GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	rates, err := rateService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the rates."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(rates)
}
