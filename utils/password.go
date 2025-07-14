package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(password string) string {
	hashedPassword,_ :=  bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CompareWithHashedPassword(hashedPassword, password string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
 return err == nil
}