package repository

import (
	"database/sql"
	"fmt"
	"model"
)

type RateRepository interface {
	// Save(rate *model.Rate) (*model.Rate, error)
	FindAll() ([]model.Rate, error)
	// DeleteAll()
}

type rateRepo struct{}

func NewRateRepository() RateRepository {
	return &rateRepo{}
}

func (*rateRepo) FindAll() ([]model.Rate, error) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorRate(err)

	// close database
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "offerId", "mark" FROM "Rate"`)
	CheckErrorRate(err)

	defer rows.Close()

	var rates []model.Rate

	for rows.Next() {
		var id string
		var offerId string
		var mark int

		err = rows.Scan(&id, &offerId, &mark)
		CheckErrorRate(err)

		rates = append(rates, model.Rate{Id: id, OfferId: offerId, Mark: mark})
	}

	return rates, nil
}

func CheckErrorRate(err error) {
	if err != nil {
		panic(err)
	}
}
