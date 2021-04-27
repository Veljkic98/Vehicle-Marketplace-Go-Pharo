package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"

	"github.com/google/uuid"
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
	//TODO: ovde ce se odvijati sva logika vezana za prikaz
	// fmt.Println(search)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorOffer(err)

	// close database
	defer db.Close()

	query := `
	SELECT "O"."id", "O"."price", "O"."publishDate", "O"."location", "V"."id", 
	"V"."make", "V"."model", "V"."date", "V"."hp", "V"."cubic"
	FROM "Offer" "O", "Vehicle" "V"
	WHERE "O"."vehicleId" = "V"."id"
	AND "O"."price" BETWEEN $1 AND $2
	AND "V"."hp" BETWEEN $3 AND $4
	AND "V"."cubic" BETWEEN $5 AND $6
	` + queryFilterMake(search.Make) + queryFilterModel(search.ModelCar) + queryFilterLocation(search.Location) +
		querySortPrice(search.PriceAscending) + querySortHP(search.HPAscending)

	rows, err := db.Query(query, search.PriceFrom, search.PriceTo, search.HPFrom, search.HPTo,
		search.CubicFrom, search.CubicTo)
	CheckErrorOffer(err)

	defer rows.Close()

	var offers []model.Offer

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

		// layout for parse string to date
		const layout = "2006-01-02"

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

	// TODO: newest

	return offers, nil
}

/*
	Find all rates by offer id.
*/
func getRatesByOffer(id string) []model.Rate {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorOffer(err)

	// close database
	defer db.Close()

	query := `
	SELECT "id", "offerId", "mark"
	FROM "Rate"
	WHERE "offerId" = $1
	`
	rows, err := db.Query(query, id)
	CheckErrorOffer(err)

	defer rows.Close()

	var rates []model.Rate

	for rows.Next() {
		var id string
		var offerId string
		var mark int

		err = rows.Scan(&id, &offerId, &mark)
		CheckErrorOffer(err)

		rates = append(rates, model.Rate{Id: id, OfferId: offerId, Mark: mark})
	}

	return rates
}

/*
	Find all comments by offer id
*/
func getCommentsByOffer(id string) []model.Comment {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorOffer(err)

	// close database
	defer db.Close()

	query := `
	SELECT "id", "offerId", "content"
	FROM "Comment"
	WHERE "offerId" = $1
	`
	rows, err := db.Query(query, id)
	CheckErrorOffer(err)

	defer rows.Close()

	var comments []model.Comment

	for rows.Next() {
		var id string
		var offerId string
		var content string

		err = rows.Scan(&id, &offerId, &content)
		CheckErrorOffer(err)

		comments = append(comments, model.Comment{Id: id, OfferId: offerId, Content: content})
	}

	return comments
}

/*
	Return part of query where
	vehicle make matches
*/
func queryFilterMake(make string) string {

	if make != "" {
		return `AND "V"."make" = '` + make + `'`
	}

	return ``
}

/*
	Return part of query where
	vehicle model matches
*/
func queryFilterModel(model string) string {

	if model != "" {
		return `AND "V"."model" = '` + model + `'`
	}

	return ``
}

/*
	Return part of query where
	offer location matches
*/
func queryFilterLocation(location string) string {

	if location != "" {
		return `AND "O"."location" = '` + location + `'`
	}

	return ``
}

/*
	Return part of query where
	sorting by price asc or desc
*/
func querySortPrice(asc bool) string {

	if asc {
		return `ORDER BY "O"."price" ASC`
	}

	return `ORDER BY "O"."price" DESC`
}

/*
	Return part of query where
	sorting by horse power asc or desc.

	Function call must go after querySortPrice().
*/
func querySortHP(asc bool) string {

	if asc {
		return `, "V"."hp" ASC`
	}

	return `, "V"."hp" DESC`

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
