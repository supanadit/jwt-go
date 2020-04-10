package ej

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strings"
)

func GetJWTFromGinHeader(c *gin.Context) (string, error) {
	var token string
	var err error = nil
	header := c.GetHeader("Authorization")
	if header != "" {
		splitAuthorization := strings.Split(header, " ")
		if len(splitAuthorization) != 0 && len(splitAuthorization) == 2 {
			token = splitAuthorization[1]
		} else {
			err = errors.New("invalid authorization header")
		}
	} else {
		err = errors.New("no authorization provided")
	}
	return token, err
}

func VerifyGinHeader(c *gin.Context) (claims *Claims, valid bool) {
	isValid := !IsUseAuthorization()
	token, err := GetJWTFromGinHeader(c)
	if err != nil {
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v \n", token.Header["alg"])
			}
			return GetJWTSecretCode(), nil
		})

		if token != nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				err = mapstructure.Decode(claims, &claims)
				if err == nil {
					isValid = true
				}
			} else {
				fmt.Printf("error while parsing token \n")
			}
		}
	}
	return claims, isValid
}
