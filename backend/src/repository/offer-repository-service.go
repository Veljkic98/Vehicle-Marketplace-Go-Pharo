package repository

import (
	"database/sql"
	"fmt"
	"model"
	"sort"
	"time"

	"github.com/google/uuid"
)

/*
	Generate query string from all parameters
	in Search object.
*/
func getQuery(search *model.Search) string {

	query := `
	SELECT "O"."id", "O"."price", "O"."publishDate", "O"."location", "V"."id", 
	"V"."make", "V"."model", "V"."date", "V"."hp", "V"."cubic"
	FROM "Offer" "O", "Vehicle" "V"
	WHERE "O"."vehicleId" = "V"."id"
	AND "O"."price" BETWEEN $1 AND $2
	AND "V"."hp" BETWEEN $3 AND $4
	AND "V"."cubic" BETWEEN $5 AND $6
	` + queryFilterMake(search.Make) + queryFilterModel(search.ModelCar) + queryFilterLocation(search.Location) +
		querySortPrice(search.Sort) + querySortHP(search.Sort)

	return query
}

/*
	Return part of query where
	vehicle make matches
*/
func queryFilterMake(make string) string {

	if make != "" {
		return `AND "V"."make" ilike '%` + make + `%'`
	}

	return ``
}

/*
	Return part of query where
	vehicle model matches
*/
func queryFilterModel(model string) string {

	if model != "" {
		return `AND "V"."model" ilike '%` + model + `%'`
	}

	return ``
}

/*
	Return part of query where
	offer location matches
*/
func queryFilterLocation(location string) string {

	if location != "" {
		return `AND "O"."location" ilike '%` + location + `%'`
	}

	return ``
}

/*
	Return part of query where
	sorting by price asc or desc
*/
func querySortPrice(sort string) string {

	if sort == "priceAscending" {
		return `ORDER BY "O"."price" ASC`
	} else if sort == "priceDescending" {
		return `ORDER BY "O"."price" DESC`
	}

	return ``
}

/*
	Return part of query where
	sorting by horse power asc or desc.

	Function call must go after querySortPrice().
*/
func querySortHP(sort string) string {

	if sort == "hpAscending" {
		return `ORDER BY "V"."hp" ASC`
	} else if sort == "hpDescending" {
		return `ORDER BY "V"."hp" DESC`
	}

	return ``

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

	var rates = []model.Rate{}

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

	// var comments []model.Comment
	var comments = []model.Comment{}

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

func CheckErrorOffer(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	Function to set price, horse power
	and cubic capacity to max, if they are
	not defined (their value is 0).
*/
func searchPreprocess(search *model.Search) {

	if search.PriceTo <= 0 {
		search.PriceTo = 99999999
	}

	if search.HPTo <= 0 {
		search.HPTo = 9999
	}

	if search.CubicTo <= 0 {
		search.CubicTo = 99999
	}
}

/*
	Remove offers whose publish date
	is not in dateFrom-dateTo range.
*/
func filterByDate(offers []model.Offer, dateFrom time.Time, dateTo time.Time) []model.Offer {

	for idx, offer := range offers {
		if !(offer.Vehicle.Date.After(dateFrom) && offer.Vehicle.Date.Before(dateTo)) {
			return append(offers[0:idx], offers[idx+1:]...)
		}
	}

	return offers
}

/*
	Sort list by publish date.

	If asc is true sort offer by newest,
	otherwise by oldest.
*/
func sortByDate(offers []model.Offer, sortStr string) []model.Offer {

	if sortStr == "newest" { // sort asc
		sort.Slice(offers, func(i, j int) bool {
			return offers[i].Date.After(offers[j].Date)
		})
	} else if sortStr == "oldest" { //sort desc
		sort.Slice(offers, func(i, j int) bool {
			return offers[i].Date.Before(offers[j].Date)
		})
	}

	return offers

}

func getOfferFromRequest(offerRequest *model.OfferRequest) *model.Offer {

	// Create vehicle
	var vehicle model.Vehicle
	vehicle.Id = uuid.New().String()
	vehicle.Make = offerRequest.Make
	vehicle.ModelCar = offerRequest.ModelCar
	vehicle.Date = offerRequest.ProductionDate
	vehicle.HP = offerRequest.HP
	vehicle.Cubic = offerRequest.Cubic

	// Empty comment and rate lists
	comments := []model.Comment{}
	rates := []model.Rate{}

	var offer model.Offer
	offer.Id = uuid.New().String()
	offer.Location = offerRequest.Location
	offer.Price = offerRequest.Price
	offer.Date = time.Now() // TODO: proveriti sta ovde vraca
	offer.Comments = comments
	offer.Rates = rates
	offer.Vehicle = vehicle

	return &offer
}
