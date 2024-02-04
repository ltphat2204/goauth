package common

import (
	"github.com/golang-jwt/jwt/v5" //sign json web token
	"os"                           // load .env
	"time"                         // get now
)

func GenerateToken(data map[string]interface{}) (string, error) {
	// Calculate expire time
	expireTime := time.Now().Add(30 * time.Minute).Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data["username"],
		"exp":      expireTime, //Expired in 30 minutes
	})

	// Sign with secret key
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
