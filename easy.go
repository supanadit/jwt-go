package ej

// Default configuration
var secretCodeJWT = "54A3E4F19C28CCA4A27E5648871A6"
var useAuthorization = true
var expiredHoursTime int64 = 1
var expiredMinutesTime int64 = 0
var expiredSecondsTime int64 = 0

// Get JWT Secret Code as a string
func GetStringJWTSecretCode() string {
	return secretCodeJWT
}

// Get JWT Secret Code as a byte
func GetJWTSecretCode() []byte {
	return []byte(secretCodeJWT)
}

// Set JWT Secret Code
func SetJWTSecretCode(secretCode string) {
	secretCodeJWT = secretCode
}

// Get status whether use authorization or not
func IsUseAuthorization() bool {
	return useAuthorization
}

// Enable authorization
func EnableAuthorization() {
	useAuthorization = true
}

// Disable the authorization
func DisableAuthorization() {
	useAuthorization = false
}

func SetExpiredTime(hour int64, minute int64, second int64) {
	expiredHoursTime = hour
	expiredMinutesTime = minute
	expiredSecondsTime = second
}
