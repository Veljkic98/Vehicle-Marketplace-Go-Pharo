package main

import (
	"database/sql"
	"fmt"
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
}

func CheckError(err error) {
	if err != nil {
		// panic(err)
		fmt.Println("DB tables are already created.")
	} else {
		fmt.Println("DB tables are created. :)")
	}
}
