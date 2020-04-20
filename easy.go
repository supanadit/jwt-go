package jwt

// secretCodeJWT is the global variable for default secret code
var secretCodeJWT = "54A3E4F19C28CCA4A27E5648871A6"

// useAuthorization is the variable for enable and disable authorization
var useAuthorization = true

// expiredHoursTime is the hours time for duration JWT token
var expiredHoursTime int64 = 1

// expiredMinutesTime is the minutes time for duration JWT token
var expiredMinutesTime int64 = 0

// expiredSecondsTime is the seconds time for duration JWT token
var expiredSecondsTime int64 = 0

// GetStringJWTSecretCode is to get JWT secret code as a string
func GetStringJWTSecretCode() string {
	return secretCodeJWT
}

// GetJWTSecretCode get JWT secret code as a byte
func GetJWTSecretCode() []byte {
	return []byte(secretCodeJWT)
}

// SetJWTSecretCode is to set JWT secret code globally
func SetJWTSecretCode(secretCode string) {
	secretCodeJWT = secretCode
}

// IsUseAuthorization is to get status whether use authorization or not
func IsUseAuthorization() bool {
	return useAuthorization
}

// EnableAuthorization is to enable authorization
func EnableAuthorization() {
	useAuthorization = true
}

// DisableAuthorization is to Disable the authorization
func DisableAuthorization() {
	useAuthorization = false
}

// SetExpiredTime is to set expired time for JWT session
func SetExpiredTime(hour int64, minute int64, second int64) {
	expiredHoursTime = hour
	expiredMinutesTime = minute
	expiredSecondsTime = second
}
