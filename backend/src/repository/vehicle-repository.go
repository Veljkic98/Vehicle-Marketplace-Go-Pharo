package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"
)

type VehicleRepository interface {
	Save(car *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
	DeleteAll()
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

	fmt.Println("---------------------------------------")
	fmt.Println(vehicle.Make)
	fmt.Println(vehicle.Id)
	fmt.Println(vehicle.Date)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	// insert to db
	insertStmt := `insert into "Vehicle"("id", "make", "model", "date", "hp", "cubic") values($1, $2, $3, $4, $5, $6)`
	_, e := db.Exec(insertStmt, vehicle.Id, vehicle.Make, vehicle.ModelCar, vehicle.Date, vehicle.HP, vehicle.Cubic)
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

	rows, err := db.Query(`SELECT "id", "make", "model", "date", "hp", "cubic" FROM "Vehicle"`)
	CheckError(err)

	defer rows.Close()

	var cars []model.Vehicle

	for rows.Next() {
		var id string
		var make string
		var modelCar string
		var date string
		var hp int
		var cubic int

		err = rows.Scan(&id, &make, &modelCar, &date, &hp, &cubic)
		CheckError(err)

		const layout = "2006-01-02"
		d, _ := time.Parse(layout, date[0:10])
		cars = append(cars, model.Vehicle{Id: id, Make: make, ModelCar: modelCar, Date: d, HP: hp, Cubic: cubic})
	}

	return cars, nil
}

func (*repo) DeleteAll() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()
	// insert to db
	insertStmt := `DELETE FROM "Vehicle"`
	_, e := db.Exec(insertStmt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
