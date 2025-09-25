package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
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

func (p publications) GetPublications() ([]models.Publication, error) {
	rows, err := p.db.Query(
		"Select id, title, content, author_id, likes, createdAt from publications",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication
		if err = rows.Scan(&publication.ID, &publication.Title, &publication.Content, &publication.AuthorID, &publication.Likes, &publication.CreatedAt); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

func (p publications) GetPublication(publicationID uint64) (models.Publication, error) {

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

// TODO - createdAt return -> verification in insomnia
func (p publications) UpdatePublication(publicationID uint64, publication models.Publication) (models.Publication, error) {

	statement, err := p.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if err != nil {
		return models.Publication{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(publication.Title, publication.Content, publicationID)
	if err != nil {
		return models.Publication{}, err
	}

	return publication, nil
}

// TODO - treatment if the publication id is not exist
func (p publications) DeletePublication(publicationID uint64) error {
	statement, err := p.db.Prepare("delete from publications where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publicationID)
	if err != nil {
		return err
	}

	return nil
}

func (p publications) GetPublicationsByUserID(userID uint64) ([]models.Publication, error) {

	rows, err := p.db.Query(`
        SELECT p.*, u.nickname
        FROM publications p
		join users u on u.id = p.author_id
        WHERE author_id = ?
    `, userID)
	if err != nil {
		fmt.Print("ERRRO AKU")
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}

	return publications, nil
}

func (p publications) LikePublication(publicationID uint64) error{
	statement, err := p.db.Prepare("update publications set likes = likes +1 where id = ?")
	if err != nil{
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(publicationID)
	if err != nil{
		return err
	}

	  rowsAffected, err := result.RowsAffected()
	  if err != nil {
		  return err
	  }
  
	  if rowsAffected == 0 {
		  return errors.New("publication not found")
	  }



	return nil
}

func (p publications) UnlikePublication(publicationID uint64) error{
	statement, err := p.db.Prepare("update publications set likes = likes -1 where id = ? and likes > 0")
	if err != nil{
		return err
	}	
	defer statement.Close()

	result, err := statement.Exec(publicationID)
	if err != nil{
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("publication not found")
	}


	return nil
}
