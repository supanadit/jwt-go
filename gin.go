package jwt

import (
	"github.com/gin-gonic/gin"
)

// GetJWTFromGinHeader is to get JWT from header which provided by gin web framework
func GetJWTFromGinHeader(c *gin.Context) (token string, err error) {
	return GetJWTFromHeader(c.GetHeader("Authorization"))
}

// VerifyGinHeader is the function that only verify gin header
func VerifyGinHeader(c *gin.Context) (string, bool, error) {
	return VerifyAndBindingGinHeader(nil, c)
}

// VerifyAndBindingGinHeader is to verify and binding the jwt model
func VerifyAndBindingGinHeader(model interface{}, c *gin.Context) (token string, isValid bool, err error) {
	token, err = GetJWTFromGinHeader(c)
	isValid, err = VerifyAndBinding(&model, token)
	return token, isValid, err
}
