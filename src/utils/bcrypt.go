package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordStr string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordStr), []byte(passwordHash))
}
