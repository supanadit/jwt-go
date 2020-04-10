package ej

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

// Default authorization model
type Authorization struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

// For JWT purpose
func (auth Authorization) Valid() error {
	panic("Unknown Error")
}

// Generate JWT Token by authorization model
func (auth Authorization) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	tokenString, err := token.SignedString(GetJWTSecretCode())
	return tokenString, err
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

// Verify the authorization
func (auth Authorization) Verify(authorizationToken string) (bool, error) {
	isValid := IsUseAuthorization()
	token, err := jwt.Parse(authorizationToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v \n", token.Header["alg"])
		}
		return GetJWTSecretCode(), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			err = mapstructure.Decode(claims, &auth)
			if err == nil {
				isValid = true
			}
		} else {
			fmt.Printf("error while parsing token \n")
		}
	}
	return isValid, err
}
