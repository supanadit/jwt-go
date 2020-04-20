package jwt

import (
	"github.com/labstack/echo/v4"
)

func GetJWTFromEchoHeader(c echo.Context) (token string, err error) {
	return GetJWTFromHeader(c.Request().Header.Get("Authorization"))
}

func VerifyEchoHeader(c echo.Context) (string, bool, error) {
	return VerifyAndBindingEchoHeader(nil, c)
}

func VerifyAndBindingEchoHeader(model interface{}, c echo.Context) (token string, isValid bool, err error) {
	token, err = GetJWTFromEchoHeader(c)
	isValid, err = VerifyAndBinding(&model, token)
	return token, isValid, err
}
