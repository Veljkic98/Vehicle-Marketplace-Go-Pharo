package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"service"
)

type OfferController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	GetAll2(response http.ResponseWriter, request *http.Request)
	Save(response http.ResponseWriter, request *http.Request)
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

func (*offerController) Save(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var offerRequest model.OfferRequest

	err := json.NewDecoder(request.Body).Decode(&offerRequest)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
		fmt.Println("Greska1")
		return
	}

	err1 := offerService.Validate(&offerRequest)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err1.Error()})
		fmt.Println("Greska2")
		return
	}

	result, err2 := offerService.Create(&offerRequest)

	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the offer."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}

func (*offerController) GetAll(response http.ResponseWriter, request *http.Request) {

	fmt.Println("*** Call GetAll Search Method ***")

	response.Header().Set("Content-Type", "application/json")

	var search model.Search
	fmt.Println(request.Body)

	err := json.NewDecoder(request.Body).Decode(&search)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
		fmt.Println("Greska1 offer cont")
		return
	}

	offers, err := offerService.FindAll(&search)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the offers."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(offers)
}

/*
	Ova metoda je samo dok se razvija klijent.

	Samo dobavljamo sve oglase.
*/
func (*offerController) GetAll2(response http.ResponseWriter, request *http.Request) {

	fmt.Println("*** Call GetAll2 Method ***")

	response.Header().Set("Content-Type", "application/json")

	offers, err := offerService.FindAll2()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting the offers."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(offers)
}
