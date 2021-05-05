package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"

	_ "github.com/lib/pq"
)

type OfferRepository interface {
	Save(offer *model.OfferRequest) (*model.Offer, error)
	FindAll(search *model.Search) ([]model.Offer, error)
	FindAll2() ([]model.Offer, error)
	// DeleteAll()
}

type offerRepo struct{}

func NewOfferRepository() OfferRepository {
	return &offerRepo{}
}

func (*offerRepo) Save(offerRequest *model.OfferRequest) (*model.Offer, error) {

	fmt.Println("------------------- adding offer --------------------")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	offer := getOfferFromRequest(offerRequest)

	///////////// vehicle offer rate comment

	// insert vehicle to db
	date := offer.Vehicle.Date.String()
	insertStmtVehicle := `insert into "Vehicle"("id", "make", "model", "date", "hp", "cubic") values($1, $2, $3, $4, $5, $6)`
	_, eVehicle := db.Exec(insertStmtVehicle, offer.Vehicle.Id, offer.Vehicle.Make, offer.Vehicle.ModelCar,
		date[0:10], offer.Vehicle.HP, offer.Vehicle.Cubic)
	CheckError(eVehicle)

	// insert offer to db
	publishDate := offer.Date.String()
	insertStmtOffer := `insert into "Offer"("id", "vehicleId", "price", "publishDate", "location") values($1, $2, $3, $4, $5)`
	_, eOffer := db.Exec(insertStmtOffer, offer.Id, offer.Vehicle.Id, offer.Price, publishDate[0:10], offer.Location)
	CheckError(eOffer)

	fmt.Println("Printamo offer da vidimo sta smo napravili")
	fmt.Println(offer)

	return offer, nil
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

func (*offerRepo) FindAll2() ([]model.Offer, error) {

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
		WHERE "O"."vehicleId" = "V"."id"`

	rows, err := db.Query(query)
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

	return offers, nil
}
