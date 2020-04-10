package ej

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// Generate JWT Token
func GenerateJWT(model jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model)
	tokenString, err := token.SignedString(GetJWTSecretCode())
	return tokenString, err
}

// Verify JWT Token
func VerifyJWT(model jwt.Claims, token string) (bool, error) {
	isValid := IsUseAuthorization()
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v \n", token.Header["alg"])
		}
		return GetJWTSecretCode(), nil
	})
	if t != nil {
		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			err = mapstructure.Decode(claims, &model)
			if err == nil {
				isValid = true
			}
		} else {
			fmt.Printf("error while parsing t \n")
		}
	}
	return isValid, err
}
