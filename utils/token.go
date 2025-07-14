package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id": userId,
	})

	signedToken,err := token.SignedString(jwtSecret)

	if err != nil {
  		return "", err
 	}
 	return signedToken, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  		return []byte(os.Getenv("JWT_SECRET")), nil
 	})
 	
	if err != nil {
  		return false, err
 	}

 	return token.Valid, nil
}