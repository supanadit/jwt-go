package ej

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetJWTFromGinHeader(c *gin.Context) (token string, err error) {
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

func VerifyGinHeader(c *gin.Context) bool {
	return VerifyAndBindingGinHeader(nil, c)
}

func VerifyAndBindingGinHeader(model interface{}, c *gin.Context) (isValid bool) {
	token, err := GetJWTFromGinHeader(c)
	if err != nil {
		isValid, err = VerifyAndBindingJWT(&model, token)
	}
	return isValid
}
