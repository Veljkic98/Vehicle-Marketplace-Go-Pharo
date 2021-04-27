package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"service"
)

type VehicleController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Save(response http.ResponseWriter, request *http.Request)
	DeleteAll(response http.ResponseWriter, request *http.Request)
}

type vehicleController struct{}

var (
	vehicleService service.VehicleService
)

func NewVehicleController(service service.VehicleService) VehicleController {
	vehicleService = service
	return &vehicleController{}
}

func (*vehicleController) GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	vehicles, err := vehicleService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the vehicles"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(vehicles)
}

func (*vehicleController) Save(response http.ResponseWriter, request *http.Request) {

	fmt.Println("-------------------save vehicle controller-----------------------")
	response.Header().Set("Content-Type", "application/json")

	var vehicle model.Vehicle

	err := json.NewDecoder(request.Body).Decode(&vehicle)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
		fmt.Println("Greska1")
		return
	}

	err1 := vehicleService.Validate(&vehicle)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err1.Error()})
		fmt.Println("Greska2")
		return
	}

	result, err2 := vehicleService.Create(&vehicle)

	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the vehicle."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}

func (*vehicleController) DeleteAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	vehicleService.DeleteAll()

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response)
}
