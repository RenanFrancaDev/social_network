package repositories

import (
	"api/src/models"
	"api/src/utils"
	"database/sql"
	"fmt"
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
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (u users) GetUsers() ([]models.User, error) {
	statement, err := u.db.Prepare(
		"SELECT id, name, nickname, email, createdAt FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
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

func (u *users) GetUserByEmail(userEmail string) (models.User, error) {

	var user models.User
	err := u.db.QueryRow("select id, password from users where email = ?", userEmail).Scan(&user.ID, &user.Password)
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

func (u *users) FollowUser(followerId uint64, userId uint64) error {
	statement, err := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) UnfollowUser(followerId uint64, userId uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) GetFollowers(userID uint64) ([]models.User, error) {

	rows, err := u.db.Query(`
	select users.id, users.name, users.nickname, users.email, users.createdAt
	from users inner join followers on users.id = followers.follower_id
	where followers.user_id = ?
	`, userID)
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

func (u *users) GetFollowing(userID uint64) ([]models.User, error) {

	rows, err := u.db.Query(`
	select users.id, users.name, users.nickname, users.email, users.createdAt
	from users inner join followers on users.id = followers.user_id
	where followers.follower_id = ?
	`, userID)
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

func (u *users) UpdatePassword(userID uint64, currentPassword string, newPassword string) error {

	var storedPassword string
	err := u.db.QueryRow("SELECT password FROM users WHERE id = ?", userID).Scan(&storedPassword)
	if err != nil {
		return err
	}

	fmt.Printf("current %s", currentPassword)
	fmt.Printf("stored %s", storedPassword)
	if err = utils.CheckPassword(storedPassword, currentPassword); err != nil {
		fmt.Print("ENTROU AQUI")
		return err
	}

	statement, err := u.db.Prepare("Update users SET password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	newHash, err := utils.Hash(newPassword)
	if err != nil {
		return err
	}

	_, err = statement.Exec(string(newHash), userID)
	if err != nil {
		return err
	}

	return nil

}
