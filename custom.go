package ej

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JWTClaims struct {
	Object interface{}
	jwt.StandardClaims
}

// Generate JWT Token
func GenerateJWT(model interface{}) (string, error) {
	r := JWTClaims{
		Object: model,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, r)
	s, e := t.SignedString(GetJWTSecretCode())
	return s, e
}

// Verify JWT Token
func VerifyJWT(model interface{}, token string) (bool, error) {
	isValid := !IsUseAuthorization()
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
			fmt.Println(err)
		}
	}
	return isValid, err
}

// Encrypt Password
func EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

// Verify between encryption password and requested password
func VerifyPassword(encryptedPassword string, password string) (bool, error) {
	isValid := false
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err == nil {
		isValid = true
	}
	return isValid, err
}
