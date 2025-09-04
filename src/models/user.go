package models

import (
	"api/src/utils"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Validations(step string) error {
	if err := u.notEmpty(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}

	return nil
}

func (u *User) notEmpty(step string) error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Nickname == "" {
		return errors.New("Nickname is required")
	}

	if step == "signup" && u.Password == "" {
		return errors.New("Password is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid format email")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nickname = strings.TrimSpace(u.Nickname)
	if step == "signup" {
		passwordHash, err := utils.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(passwordHash)
	}

	return nil
}
