package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type publications struct {
	db *sql.DB
}

func NewUsersRepositoryPulications(db *sql.DB) *publications {
	return &publications{db}
}

func (p publications) Create(userID uint64, publication models.Publication) (uint64, error) {

	statement, err := p.db.Prepare(
		"insert into publications (title, content,author_id) values (?, ?, ?)",
	)
	if err != nil {

		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, userID)
	if err != nil {
		fmt.Printf("ERRORRES %v", err)
		return 0, err
	}

	fmt.Printf("Result %v", result)

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastId), nil

}
