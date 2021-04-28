package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"service"
)

type RateController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Save(response http.ResponseWriter, request *http.Request)
	// DeleteAll(response http.ResponseWriter, request *http.Request)
}

type rateController struct{}

var (
	rateService service.RateService
)

func NewRateController(service service.RateService) RateController {
	rateService = service
	return &rateController{}
}

func (*rateController) Save(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var rate model.Rate

	err := json.NewDecoder(request.Body).Decode(&rate)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
		fmt.Println("Error 1 rate")
		return
	}

	err1 := rateService.Validate(&rate)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err1.Error()})
		fmt.Println("Error 2 rate")
		return
	}

	result, err2 := rateService.Create(&rate)

	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the rate."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
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
