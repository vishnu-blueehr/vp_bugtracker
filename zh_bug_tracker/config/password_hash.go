package config

import (
	"fmt"

	// this package is used for Encrtypting password

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	Hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println("Hashed password:   ", Hash)
	return string(Hash), err
}

func CheckPasswordHash(password, hash string) bool {
	fmt.Println("password:   ", password)
	fmt.Println("hashed password:   ", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
