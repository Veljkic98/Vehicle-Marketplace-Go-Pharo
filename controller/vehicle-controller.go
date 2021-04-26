package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"service"
)

type VehicleController interface {
	GetAllVehicles(response http.ResponseWriter, request *http.Request)
	SaveVehicle(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

var (
	vehicleService service.VehicleService
)

func NewVehicleController(service service.VehicleService) VehicleController {
	vehicleService = service
	return &controller{}
}

func (*controller) GetAllVehicles(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	vehicles, err := vehicleService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the vehicles"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(vehicles)
}

func (*controller) SaveVehicle(response http.ResponseWriter, request *http.Request) {

	fmt.Printf("------------------------------------------")
	response.Header().Set("Content-Type", "application/json")

	var vehicle model.Vehicle

	err := json.NewDecoder(request.Body).Decode(&vehicle)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
		return
	}

	err1 := vehicleService.Validate(&vehicle)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err1.Error()})
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
