package repository

import (
	"model"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
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

	// // connection string
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// // open database
	// db, err := sql.Open("postgres", psqlconn)
	// CheckErrorOffer(err)

	// // close database
	// defer db.Close()

	// rows, err := db.Query(`SELECT "id", "price", "publishDate", "location" FROM "Offer"`)
	// CheckErrorOffer(err)

	// defer rows.Close()

	var offers []model.Offer

	// for rows.Next() {
	// 	var id string
	// 	// var vehicleId string
	// 	var price int
	// 	var publishDate string
	// 	var location string

	// 	err = rows.Scan(&id, &price, &publishDate, &location)
	// 	CheckErrorOffer(err)

	// 	const layout = "2006-01-02"
	// 	d, _ := time.Parse(layout, publishDate[0:10])
	// 	offers = append(offers, model.Offer{Id: id, Price: price, Date: d, Location: location})
	// }

	/////////////////////////////////////////////////////////////////////////
	// const layout = "2006-01-02"
	// d, _ := time.Parse(layout, "2020-05-05")

	var offer model.Offer
	offer = getOffer()

	offers = append(offers, offer)

	/////////////////////////////////////////////////////////////////////////

	return offers, nil
}

func CheckErrorOffer(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	manually create and return offer
*/
func getOffer() model.Offer {
	const layout = "2006-01-02"
	d, _ := time.Parse(layout, "2020-05-05")

	var offer model.Offer
	offer.Id = uuid.New().String()
	offer.Location = "Novi Sad"
	offer.Date = d
	offer.Price = 20000

	var vehicle model.Vehicle

	vehicle.Id = "1"
	vehicle.Make = "BMW"
	vehicle.ModelCar = "320"
	vehicle.HP = 150
	vehicle.Cubic = 2000

	vehicle.Date = d

	// add vehicle to offer
	offer.Vehicle = vehicle

	var comment1 model.Comment
	comment1.Id = uuid.New().String()
	comment1.OfferId = offer.Id
	comment1.Content = "com 1"

	// add comment to offer
	offer.Comments = append(offer.Comments, comment1)

	var rate1 model.Rate
	rate1.Id = uuid.New().String()
	rate1.OfferId = offer.Id
	rate1.Mark = 3

	// add rate to offer
	offer.Rates = append(offer.Rates, rate1)

	return offer
}
