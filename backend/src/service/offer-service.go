package service

import (
	"errors"
	"model"
	"repository"
	"time"
)

type OfferService interface {
	Validate(offerRequest *model.OfferRequest) error
	Create(offer *model.OfferRequest) (*model.Offer, error)
	FindAll(search *model.Search) ([]model.Offer, error)
	// DeleteAll()
}

type offerService struct{}

var (
	offerRepo repository.OfferRepository
)

func NewOfferService(repo repository.OfferRepository) OfferService {
	offerRepo = repo
	return &offerService{}
}

func (*offerService) Create(offerRequest *model.OfferRequest) (*model.Offer, error) {

	return offerRepo.Save(offerRequest)
}

func (*offerService) Validate(offerRequest *model.OfferRequest) error {

	if offerRequest == nil {
		err := errors.New("the offer request is empty.")
		return err
	}

	if offerRequest.Price == 0 {
		err := errors.New("the vehicle price is empty.")
		return err
	}

	if offerRequest.Location == "" {
		err := errors.New("the location is empty.")
		return err
	}

	if offerRequest.Make == "" {
		err := errors.New("the vehicle make is empty.")
		return err
	}

	if offerRequest.ModelCar == "" {
		err := errors.New("the vehicle model is empty.")
		return err
	}

	if offerRequest.HP == 0 {
		err := errors.New("the vehicle HP is empty.")
		return err
	}

	if offerRequest.Cubic == 0 {
		err := errors.New("the vehicle cubic capacity is empty.")
		return err
	}

	var zeroTime time.Time

	if offerRequest.ProductionDate == zeroTime {
		err := errors.New("the vehicle date is empty.")
		return err
	}

	return nil
}

func (*offerService) FindAll(search *model.Search) ([]model.Offer, error) {

	return offerRepo.FindAll(search)
}
