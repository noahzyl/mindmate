package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash string by bcrypt
func HashStringWithBcrypt(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckBcryptHash(string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(string))
	return err == nil
}
