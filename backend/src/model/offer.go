package model

import "time"

type Offer struct {
	Id        string    `json:"id"`
	VehicleId string    `json:"vehicleId"`
	Price     int       `json:"price"`
	Date      time.Time `json:"date"` // publish date
	Location  string    `json:"location"`
}
