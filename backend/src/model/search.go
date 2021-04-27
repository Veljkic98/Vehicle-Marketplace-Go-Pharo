package model

import "time"

type Search struct {
	// Filter
	Make      string    `json:"make"`
	ModelCar  string    `json:"model"`
	Location  string    `json:"location"`
	PriceFrom int       `json:"priceFrom"`
	PriceTo   int       `json:"priceTo"`
	DateFrom  time.Time `json:"dateFrom"`
	DateTo    time.Time `json:"dateTo"`
	HPFrom    int       `json:"hpFrom"`
	HPTo      int       `json:"hpTo"`
	CubicFrom int       `json:"cubicFrom"`
	CubicTo   int       `json:"cubicTo"`

	// Sort
	Newest         bool `json:"newest"` // looking offer publish date
	PriceAscending bool `json:"priceAscending"`
	HPAscending    bool `json:"hpAscending"`
}
