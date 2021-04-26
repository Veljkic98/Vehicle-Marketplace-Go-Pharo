package model

// "github.com/jinzhu/gorm"

type Vehicle struct {
	// gorm.Model
	// Year      int
	Id   string `json:"id"`
	Make string `json:"make"`
	// ModelName string
	// DriverID   int
}
