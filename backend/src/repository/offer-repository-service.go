package repository

import (
	"database/sql"
	"fmt"
	"model"
)

/*
	Generate query string from all parameters
	in Search object.
*/
func getQuery(search *model.Search) string {

	fmt.Println(search.PriceTo)
	searchPreprocess(search)
	fmt.Println(search.PriceTo)

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

	return query
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

	if search.PriceTo == 0 {
		search.PriceTo = 99999999
	}

	if search.HPTo == 0 {
		search.HPTo = 9999
	}

	if search.CubicTo == 0 {
		search.CubicTo = 9999
	}
}
