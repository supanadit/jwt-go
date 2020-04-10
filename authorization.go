package ej

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type Authorization struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func (auth Authorization) Valid() error {
	panic("Unknown Error")
}

func (auth Authorization) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	tokenString, err := token.SignedString(GetJWTSecretCode())
	return tokenString, err
}

func (auth Authorization) GetPasswordEncryption() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(auth.Password), 12)
}

func (auth Authorization) GetStringPasswordEncryption() (string, error) {
	encryption, err := auth.GetPasswordEncryption()
	return string(encryption), err
}

func (auth Authorization) Verify(authorizationToken string) (bool, error) {
	isValid := GetAuthorizationStatus()
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
