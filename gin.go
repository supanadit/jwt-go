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

func VerifyGinHeader(model Claims, c *gin.Context) bool {
	isValid := !IsUseAuthorization()
	token, err := GetJWTFromGinHeader(c)
	if err != nil {
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
					fmt.Println(err)
				}
				isValid = true
			} else {
				fmt.Println(err)
			}
		}
	}
	return isValid
}
