package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Validations() error {
	if err := u.notEmpty(); err != nil {
		return err
	}

	u.cleanSpaces()

	return nil
}

func (u *User) notEmpty() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Nickname == "" {
		return errors.New("Nickname is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	return nil
}

func (u *User) cleanSpaces() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nickname = strings.TrimSpace(u.Nickname)
}
