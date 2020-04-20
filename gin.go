package jwt

import (
	"github.com/gin-gonic/gin"
)

func GetJWTFromGinHeader(c *gin.Context) (token string, err error) {
	return GetJWTFromHeader(c.GetHeader("Authorization"))
}

func VerifyGinHeader(c *gin.Context) (string, bool, error) {
	return VerifyAndBindingGinHeader(nil, c)
}

func VerifyAndBindingGinHeader(model interface{}, c *gin.Context) (token string, isValid bool, err error) {
	token, err = GetJWTFromGinHeader(c)
	isValid, err = VerifyAndBinding(&model, token)
	return token, isValid, err
}
