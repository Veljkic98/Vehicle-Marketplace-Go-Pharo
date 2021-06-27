package repository

import (
	"database/sql"
	"fmt"
	"model"

	_ "github.com/lib/pq"
)

type CommentRepository interface {
	Save(comment *model.Comment) (*model.Comment, error)
	FindAll() ([]model.Comment, error)
	// DeleteAll()
}

type commentRepo struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepo{}
}

func (*commentRepo) Save(comment *model.Comment) (*model.Comment, error) {

	fmt.Println("*** Adding comment ***")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	// insert to db
	insertStmt := `insert into "Comment"("id", "offerId", "content") values($1, $2, $3)`
	_, e := db.Exec(insertStmt, comment.Id, comment.OfferId, comment.Content)
	CheckError(e)

	return comment, nil
}

func (*commentRepo) FindAll() ([]model.Comment, error) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckErrorComment(err)

	// close database
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "offerId", "content" FROM "Comment"`)
	CheckErrorComment(err)

	defer rows.Close()

	var comments []model.Comment

	for rows.Next() {
		var id string
		var offerId string
		var content string

		err = rows.Scan(&id, &offerId, &content)
		CheckErrorComment(err)

		comments = append(comments, model.Comment{Id: id, OfferId: offerId, Content: content})
	}

	return comments, nil
}

func CheckErrorComment(err error) {
	if err != nil {
		panic(err)
	}
}
