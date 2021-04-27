package model

import "time"

// "github.com/jinzhu/gorm"

type Vehicle struct {
	// gorm.Model
	// Year      int
	Id       string    `json:"id"`
	Make     string    `json:"make"`
	ModelCar string    `json:"modelCar"`
	Date     time.Time `json:"date"` // production date
	HP       int       `json:"hp"`
	Cubic    int       `json:"cubic"`
	// ModelName string
	// DriverID   int
}
