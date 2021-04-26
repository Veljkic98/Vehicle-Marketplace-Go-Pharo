package service

import (
	"errors"
	"model"
	"repository"

	"github.com/google/uuid"
)

type VehicleService interface {
	Validate(vehicle *model.Vehicle) error
	Create(vehicle *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
}

type service struct{}

var (
	vehicleRepo repository.VehicleRepository
)

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	vehicleRepo = repo
	return &service{}
}

func (*service) Validate(vehicle *model.Vehicle) error {

	if vehicle == nil {
		err := errors.New("The vehicle is empty.")
		return err
	}

	if vehicle.Make == "" {
		err := errors.New("The make is empty.")
		return err
	}

	return nil
}

func (*service) Create(vehicle *model.Vehicle) (*model.Vehicle, error) {

	vehicle.Id = uuid.New().String()

	return vehicleRepo.Save(vehicle)
}

func (*service) FindAll() ([]model.Vehicle, error) {

	return vehicleRepo.FindAll()
}
