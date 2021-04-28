package main

import (
	"database/sql"
	"fmt"
	"model"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "veljkosql"
	dbname   = "ntp"
)

/*
	Method to create all tables.
*/
func createTablesDB() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	// insert to db
	insertStmt := `
	CREATE TABLE "Vehicle"(
		"id" TEXT NOT NULL,
		"make" TEXT NOT NULL,
		"model" TEXT NOT NULL,
		"date" TEXT NOT NULL,
		"hp" INTEGER NOT NULL,
		"cubic" INTEGER NOT NULL
	);
	ALTER TABLE
		"Vehicle" ADD PRIMARY KEY("id");
	CREATE TABLE "Offer"(
		"id" TEXT NOT NULL,
		"vehicleId" TEXT NOT NULL,
		"price" INTEGER NOT NULL,
		"publishDate" TEXT NOT NULL,
		"location" TEXT NOT NULL
	);
	ALTER TABLE
		"Offer" ADD PRIMARY KEY("id");
	CREATE TABLE "Comment"(
		"id" TEXT NOT NULL,
		"offerId" TEXT NOT NULL,
		"content" TEXT NOT NULL
	);
	ALTER TABLE
		"Comment" ADD PRIMARY KEY("id");
	CREATE TABLE "Rate"(
		"id" TEXT NOT NULL,
		"offerId" TEXT NOT NULL,
		"mark" INTEGER NOT NULL
	);
	ALTER TABLE
		"Rate" ADD PRIMARY KEY("id");
	ALTER TABLE
		"Offer" ADD CONSTRAINT "offer_vehicleid_foreign" FOREIGN KEY("vehicleId") REFERENCES "Vehicle"("id");
	ALTER TABLE
		"Comment" ADD CONSTRAINT "comment_offerid_foreign" FOREIGN KEY("offerId") REFERENCES "Offer"("id");
	ALTER TABLE
		"Rate" ADD CONSTRAINT "rate_offerid_foreign" FOREIGN KEY("offerId") REFERENCES "Offer"("id");
	`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	fmt.Println("---------- DB tables are created ----------")
}

/*
	Create one entity for all tables.
*/
func createAllInit() {

	var offer model.Offer
	offer = getOffer()

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	var insertStmt string

	// insert vehicle into db
	date := offer.Vehicle.Date.String()
	insertStmt = `insert into "Vehicle"("id", "make", "model", "date", "hp", "cubic") values($1, $2, $3, $4, $5, $6)`
	_, e1 := db.Exec(insertStmt, offer.Vehicle.Id, offer.Vehicle.Make, offer.Vehicle.ModelCar,
		date[0:10], offer.Vehicle.HP, offer.Vehicle.Cubic)
	CheckError(e1)

	// insert offer into db
	publishDate := offer.Date.String()
	insertStmt = `insert into "Offer"("id", "vehicleId", "price", "publishDate", "location") values($1, $2, $3, $4, $5)`
	_, e2 := db.Exec(insertStmt, offer.Id, offer.Vehicle.Id, offer.Price,
		publishDate[0:10], offer.Location)
	CheckError(e2)

	// insert rate into db
	insertStmt = `insert into "Rate"("id", "offerId", "mark") values($1, $2, $3)`
	_, e3 := db.Exec(insertStmt, offer.Rates[0].Id, offer.Id, offer.Rates[0].Mark)
	CheckError(e3)

	// insert comment into db
	insertStmt = `insert into "Comment"("id", "offerId", "content") values($1, $2, $3)`
	_, e4 := db.Exec(insertStmt, offer.Comments[0].Id, offer.Id, offer.Comments[0].Content)
	CheckError(e4)

	fmt.Println("---------- Entities added to db ----------")
}

/*
	Delete all entities from all tables
*/
func deleteAll() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	var insertStmt string

	// delete all from rate
	insertStmt = `DELETE FROM "Rate"`
	_, e1 := db.Exec(insertStmt)
	CheckError(e1)

	// delete all from comment
	insertStmt = `DELETE FROM "Comment"`
	_, e2 := db.Exec(insertStmt)
	CheckError(e2)

	// delete all from offer
	insertStmt = `DELETE FROM "Offer"`
	_, e4 := db.Exec(insertStmt)
	CheckError(e4)

	// delete all from vehicle
	insertStmt = `DELETE FROM "Vehicle"`
	_, e3 := db.Exec(insertStmt)
	CheckError(e3)

}

/*
	manually create and return offer
*/
func getOffer() model.Offer {
	const layout = "2006-01-02"
	d, _ := time.Parse(layout, "2018-05-05")

	var offer model.Offer
	offer.Id = uuid.New().String()
	offer.Location = "Novi Sad"
	offer.Date = d
	offer.Price = 16000

	var vehicle model.Vehicle

	vehicle.Id = uuid.New().String()
	vehicle.Make = "VW"
	vehicle.ModelCar = "Golf"
	vehicle.HP = 116
	vehicle.Cubic = 1600

	vehicle.Date = d

	// add vehicle to offer
	offer.Vehicle = vehicle

	var comment1 model.Comment
	comment1.Id = uuid.New().String()
	comment1.OfferId = offer.Id
	comment1.Content = "Golfic grmi"

	// add comment to offer
	offer.Comments = append(offer.Comments, comment1)

	var rate1 model.Rate
	rate1.Id = uuid.New().String()
	rate1.OfferId = offer.Id
	rate1.Mark = 4

	// add rate to offer
	offer.Rates = append(offer.Rates, rate1)

	return offer
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
