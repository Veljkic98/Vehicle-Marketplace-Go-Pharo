package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetData(db *sql.DB) {
	// rows, err := db.Query(`SELECT "name", "roll" FROM "student"`)
	rows, err := db.Query(`SELECT "id", "make" FROM "Vehicle"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var make string

		err = rows.Scan(&id, &make)
		CheckError(err)

		fmt.Println(id, make)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
