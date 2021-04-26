package repository

import (
	"database/sql"
	"fmt"
	"model"
)

type VehicleRepository interface {
	Save(car *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
}

type repo struct{}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "veljkosql"
	dbname   = "ntp"
)

func NewVehicleRepository() VehicleRepository {
	return &repo{}
}

func (*repo) Save(vehicle *model.Vehicle) (*model.Vehicle, error) {

	// TODO: save
	fmt.Println(vehicle.Make)
	fmt.Println(vehicle.Id)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	// insert to db
	insertStmt := `insert into "Vehicle"("id", "make") values($1, $2)`
	_, e := db.Exec(insertStmt, vehicle.Id, vehicle.Make)
	CheckError(e)

	return vehicle, nil
}

func (*repo) FindAll() ([]model.Vehicle, error) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "make" FROM "Vehicle"`)
	CheckError(err)

	defer rows.Close()

	var cars []model.Vehicle

	for rows.Next() {
		var id string
		var make string

		err = rows.Scan(&id, &make)
		CheckError(err)

		// fmt.Println(id, make)
		cars = append(cars, model.Vehicle{Id: id, Make: make})
	}

	// cars = append(cars, model.Vehicle{Id: "111111", Make: "BMW"})

	return cars, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
