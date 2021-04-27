package service

import (
	"errors"
	"model"
	"repository"
)

type OfferService interface {
	Validate(offer *model.Offer) error
	// Create(offer *model.Offer) (*model.Offer, error)
	FindAll() ([]model.Offer, error)
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

func (*offerService) Validate(offer *model.Offer) error {

	if offer == nil {
		err := errors.New("the offer is empty.")
		return err
	}

	if offer.VehicleId == "" {
		err := errors.New("the vehicle id is empty.")
		return err
	}

	return nil
}

func (*offerService) FindAll() ([]model.Offer, error) {

	return offerRepo.FindAll()
}
