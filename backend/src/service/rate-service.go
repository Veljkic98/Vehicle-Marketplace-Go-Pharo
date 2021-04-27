package service

import (
	"errors"
	"model"
	"repository"
)

type RateService interface {
	Validate(rate *model.Rate) error
	// Create(rate *model.Rate) (*model.Rate, error)
	FindAll() ([]model.Rate, error)
	// DeleteAll()
}

type rateService struct{}

var (
	rateRepo repository.RateRepository
)

func NewRateService(repo repository.RateRepository) RateService {
	rateRepo = repo
	return &rateService{}
}

func (*rateService) Validate(rate *model.Rate) error {

	if rate == nil {
		err := errors.New("the rate is empty.")
		return err
	}

	if rate.OfferId == "" {
		err := errors.New("the offer id is empty.")
		return err
	}

	return nil
}

func (*rateService) FindAll() ([]model.Rate, error) {

	return rateRepo.FindAll()
}
