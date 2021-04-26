package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "veljkosql"
	dbname   = "ntp"
)

func AddVehicle() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// dynamic
	// TODO: zakucane su vrednosti
	insertDynStmt := `insert into "Vehicle"("id", "make") values($1, $2)`
	_, e := db.Exec(insertDynStmt, 1, "BMW")
	CheckError(e)

	GetData(db)
}
