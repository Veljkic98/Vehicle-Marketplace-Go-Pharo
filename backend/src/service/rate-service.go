package service

import (
	"errors"
	"model"
	"repository"

	"github.com/google/uuid"
)

type RateService interface {
	Validate(rate *model.Rate) error
	Create(rate *model.Rate) (*model.Rate, error)
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

func (*rateService) Create(rate *model.Rate) (*model.Rate, error) {

	rate.Id = uuid.New().String()

	return rateRepo.Save(rate)
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

	if rate.Mark < 0 || rate.Mark > 5 {
		err := errors.New("Mark must be in range 1-5.")
		return err
	}

	return nil
}

func (*rateService) FindAll() ([]model.Rate, error) {

	return rateRepo.FindAll()
}
