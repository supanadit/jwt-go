package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CustomClaims struct {
	Object interface{}
	jwt.StandardClaims
}

// Generate JWT Token
func GenerateJWT(model interface{}) (s string, e error) {
	return GenerateJWTAndSetExpiredTime(model, expiredHoursTime, expiredMinutesTime, expiredSecondsTime)
}

// Generate JWT Token
func GenerateJWTAndSetExpiredTime(model interface{}, hours int64, minutes int64, seconds int64) (s string, e error) {
	if hours != 0 || minutes != 0 || seconds != 0 {
		r := CustomClaims{
			Object: model,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(hours) +
					time.Minute*time.Duration(minutes) +
					time.Second*time.Duration(seconds)).Unix(),
			},
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, r)
		s, e = t.SignedString(GetJWTSecretCode())
	} else {
		e = errors.New("expired time must be at least 1 second")
	}
	return s, e
}

// Verify JWT Token
func VerifyJWT(token string) (bool, error) {
	return VerifyAndBindingJWT(nil, token)
}

func VerifyAndBindingJWT(model interface{}, token string) (bool, error) {
	isValid := !IsUseAuthorization()
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecretCode(), nil
	})
	if t != nil {
		if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
			if model != nil {
				err = mapstructure.Decode(claims.Object, &model)
			}
			isValid = true
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
