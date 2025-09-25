package utils

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 12).Unix() //token expire time
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	//convert token in JSON to check values
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("[utils/token] Invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")

	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

// Verification signature method
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	//TODO - ENTENDERRR o !ok
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signature method unexpected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	//mapClaims return interface and UserId turned a float, ites necessery extract and convert to uint 64
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := permissions["userID"].(float64)
		if !ok {
			return 0, errors.New("userID is not a float")
		}
		return uint64(userID), nil
	}

	return 0, errors.New("invalid token")
}
