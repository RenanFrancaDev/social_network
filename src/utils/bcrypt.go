package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordHash string, passwordStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordStr))
}
