package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Insert an user into DB
func (u users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (name, nickname, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		log.Print("[repository] [msg: Error in DB prepare]")
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		log.Print("[repository] [msg: Error in Exec to DB]")
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Print("[repository] [msg: Error to catch last ID]")
		return 0, err
	}

	return uint64(lastId), nil
}

func (u users) GetUsers() ([]models.User, error) {
	statement, err := u.db.Prepare(
		"SELECT id, name, nickname, email, createdAt FROM users",
	)
	if err != nil {
		log.Print("[repository] [msg: Error in DB prepare]")
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Print("[repository] [msg: Error in DB Query]")
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			log.Print("[repository] [msg: Error in looping over rows]")
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (u users) SearchUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	rows, err := u.db.Query("select id, name, nickname, email, createdAt from users where nickname LIKE ? or name LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *users) GetUser(id uint64) (models.User, error) {

	var user models.User

	err := u.db.QueryRow("select id, name, nickname, email, createdAt from users where id = ?", id).Scan(
		&user.ID,
		&user.Name,
		&user.Nickname,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}

func (u *users) UpdateUser(id uint64, user models.User) (models.User, error) {

	statement, err := u.db.Prepare("update users set name = ?, nickname = ?, email = ? where id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *users) DeleteUser(id uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
