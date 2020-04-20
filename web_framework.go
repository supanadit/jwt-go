package jwt

import (
	"errors"
	"strings"
)

func GetJWTFromHeader(header string) (token string, err error) {
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

func VerifyAndBinding(model interface{}, t string) (bool, error) {
	isValid, err := VerifyAndBindingJWT(&model, t)
	if err != nil {
		if !IsUseAuthorization() {
			err = nil
			isValid = !IsUseAuthorization()
		}
	}
	return isValid, err
}
