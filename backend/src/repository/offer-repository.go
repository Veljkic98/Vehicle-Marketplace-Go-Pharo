package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"
)

type OfferRepository interface {
	// Save(offer *model.Offer) (*model.Offer, error)
	FindAll() ([]model.Offer, error)
	// DeleteAll()
}

type offerRepo struct{}

func NewOfferRepository() OfferRepository {
	return &offerRepo{}
}

func (*offerRepo) FindAll() ([]model.Offer, error) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorOffer(err)

	// close database
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "vehicleId", "price", "publishDate", "location" FROM "Offer"`)
	CheckErrorOffer(err)

	defer rows.Close()

	var offers []model.Offer

	for rows.Next() {
		var id string
		var vehicleId string
		var price int
		var publishDate string
		var location string

		err = rows.Scan(&id, &vehicleId, &price, &publishDate, &location)
		CheckErrorOffer(err)

		const layout = "2006-01-02"
		d, _ := time.Parse(layout, publishDate[0:10])
		offers = append(offers, model.Offer{Id: id, VehicleId: vehicleId, Price: price, Date: d, Location: location})
	}

	return offers, nil
}

func CheckErrorOffer(err error) {
	if err != nil {
		panic(err)
	}
}
