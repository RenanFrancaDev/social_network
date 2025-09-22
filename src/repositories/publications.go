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

func (p publications) GetPublications() ([]models.Publication, error){
	rows, err := p.db.Query(
		"Select id, title, content, author_id, likes, createdAt from publications",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next(){
		var publication models.Publication
		if err = rows.Scan(&publication.ID, &publication.Title, &publication.Content, &publication.AuthorID, &publication.Likes, &publication.CreatedAt); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

func (p publications) GetPublication(publicationID uint64) (models.Publication, error){
	
	var publication models.Publication 

	err := p.db.QueryRow(
		"Select id, title, content, author_id, likes, createdAt from publications where id = ? ", publicationID).Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
		)
	if err != nil {
		return models.Publication{}, err
	}

	return publication, nil


}