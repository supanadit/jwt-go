package jwt

import (
	"github.com/labstack/echo/v4"
)

// GetJWTFromEchoHeader get JWT from header which provided by echo web framework
func GetJWTFromEchoHeader(c echo.Context) (token string, err error) {
	return GetJWTFromHeader(c.Request().Header.Get("Authorization"))
}

// VerifyEchoHeader only verify echo header
func VerifyEchoHeader(c echo.Context) (string, bool, error) {
	return VerifyAndBindingEchoHeader(nil, c)
}

// VerifyAndBindingEchoHeader verify and binding the jwt model
func VerifyAndBindingEchoHeader(model interface{}, c echo.Context) (token string, isValid bool, err error) {
	token, err = GetJWTFromEchoHeader(c)
	isValid, err = VerifyAndBinding(&model, token)
	return token, isValid, err
}
