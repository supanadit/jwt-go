package ej

import (
	"golang.org/x/crypto/bcrypt"
)

// Default authorization model
type Authorization struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// For JWT purpose
func (auth Authorization) Valid() error {
	panic("Unknown Error")
}

// Generate JWT Token by authorization model
func (auth Authorization) GenerateJWT() (string, error) {
	return GenerateJWT(auth)
}

// Create password encryption using bcrypt
func (auth Authorization) GetPasswordEncryption() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(auth.Password), 12)
}

// Create password encryption using bcrypt and string as the result
func (auth Authorization) GetStringPasswordEncryption() (string, error) {
	encryption, err := auth.GetPasswordEncryption()
	return string(encryption), err
}

// VerifyJWT the authorization
func (auth Authorization) VerifyJWT(token string) (bool, error) {
	return VerifyJWT(auth, token)
}
