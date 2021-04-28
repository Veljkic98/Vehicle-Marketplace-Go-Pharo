package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"

	_ "github.com/lib/pq"
)

type OfferRepository interface {
	// Save(offer *model.Offer) (*model.Offer, error)
	FindAll(search *model.Search) ([]model.Offer, error)
	// DeleteAll()
}

type offerRepo struct{}

func NewOfferRepository() OfferRepository {
	return &offerRepo{}
}

func (*offerRepo) FindAll(search *model.Search) ([]model.Offer, error) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorOffer(err)

	// close database
	defer db.Close()

	searchPreprocess(search)

	query := getQuery(search)

	rows, err := db.Query(query, search.PriceFrom, search.PriceTo, search.HPFrom, search.HPTo,
		search.CubicFrom, search.CubicTo)
	CheckErrorOffer(err)

	defer rows.Close()

	offers := []model.Offer{}

	// layout for parse string to date
	const layout = "2006-01-02"

	for rows.Next() {
		var id string
		var price int
		var publishDate string
		var location string
		var vehicleId string
		var make string
		var modelCar string
		var date string
		var hp int
		var cubic int

		err = rows.Scan(&id, &price, &publishDate, &location,
			&vehicleId, &make, &modelCar, &date, &hp, &cubic)
		CheckErrorOffer(err)

		// Create vehicle
		var vehicle model.Vehicle
		vehicle.Id = vehicleId
		vehicle.ModelCar = modelCar
		vehicle.Make = make
		d1, _ := time.Parse(layout, date[0:10])
		vehicle.Date = d1
		vehicle.HP = hp
		vehicle.Cubic = cubic

		// create rates
		rates := getRatesByOffer(id)

		// create comments
		comments := getCommentsByOffer(id)

		d2, _ := time.Parse(layout, publishDate[0:10])
		offers = append(offers, model.Offer{Id: id, Price: price, Date: d2,
			Location: location, Vehicle: vehicle, Rates: rates, Comments: comments})
	}

	offers = filterByDate(offers, search.DateFrom, search.DateTo)
	offers = sortByDate(offers, search.Sort)

	return offers, nil
}
