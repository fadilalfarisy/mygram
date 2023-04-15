package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 8
	password := []byte(pass)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(hashedPass, pass []byte) bool {
	hash, password := []byte(hashedPass), []byte(pass)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
