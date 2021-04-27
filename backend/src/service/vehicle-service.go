package service

import (
	"errors"
	"model"
	"repository"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type VehicleService interface {
	Validate(vehicle *model.Vehicle) error
	Create(vehicle *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
	DeleteAll()
}

type vehicleService struct{}

var (
	vehicleRepo repository.VehicleRepository
)

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	vehicleRepo = repo
	return &vehicleService{}
}

func (*vehicleService) Validate(vehicle *model.Vehicle) error {

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

func (*vehicleService) Create(vehicle *model.Vehicle) (*model.Vehicle, error) {

	vehicle.Id = uuid.New().String()

	return vehicleRepo.Save(vehicle)
}

func (*vehicleService) FindAll() ([]model.Vehicle, error) {

	return vehicleRepo.FindAll()
}

func (*vehicleService) DeleteAll() {

	vehicleRepo.DeleteAll()
}
