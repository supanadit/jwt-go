package jwt

// Authorization is the default authorization model
type Authorization struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// GenerateJWT is to generate JWT token by authorization model
func (auth Authorization) GenerateJWT() (string, error) {
	return GenerateJWT(auth)
}

// GenerateJWTAndSetExpiredTime generate JWT token by authorization model also set the expired time manually
func (auth Authorization) GenerateJWTAndSetExpiredTime(hour int64, minute int64, seconds int64) (string, error) {
	return GenerateJWTAndSetExpiredTime(auth, hour, minute, seconds)
}

// GetPasswordEncryption is to create password encryption using bcrypt
func (auth Authorization) GetPasswordEncryption() ([]byte, error) {
	return EncryptPassword(auth.Password)
}

// GetStringPasswordEncryption is to create password encryption using bcrypt and string as the result
func (auth Authorization) GetStringPasswordEncryption() (string, error) {
	encryption, err := auth.GetPasswordEncryption()
	return string(encryption), err
}

// VerifyJWT is the simply way to check authorization
func (auth Authorization) VerifyJWT(token string) (bool, error) {
	return VerifyJWT(token)
}

// VerifyPassword is the fastest way to verify password
func (auth Authorization) VerifyPassword(password string) (bool, error) {
	ep, err := auth.GetStringPasswordEncryption()
	if err != nil {
		return false, err
	}
	return VerifyPassword(ep, password)
}
