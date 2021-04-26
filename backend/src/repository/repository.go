package repository

import "model"

type VehicleRepository interface {
	Save(car *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
}

type repo struct{}

func NewVehicleRepository() VehicleRepository {
	return &repo{}
}

func (*repo) Save(vehicle *model.Vehicle) (*model.Vehicle, error) {

	// TODO: save
	return vehicle, nil
}

func (*repo) FindAll() ([]model.Vehicle, error) {

	// TODO: find all from db
	var cars []model.Vehicle

	cars = append(cars, model.Vehicle{Id: "111111", Make: "BMW"})
	cars = append(cars, model.Vehicle{Id: "222222", Make: "Merc"})

	return cars, nil
}
