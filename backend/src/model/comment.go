package model

type Comment struct {
	Id      string `json:"id"`
	OfferId string `json:"offerId"`
	Content string `json:"content"`
}
