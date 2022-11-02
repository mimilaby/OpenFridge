// Package utils contains all the helper functions
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword compares hashed passwords
func ComparePassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
