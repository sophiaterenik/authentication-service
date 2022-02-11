package account

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

var d = Database{}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Database struct {
	Accounts map[string]string
}

func (s *Database) Initialize() {
	s.Accounts = make(map[string]string)
}

func Store(u string, hash string) {
	if d.Accounts == nil {
		d.Initialize()
	}
	d.Accounts[u] = hash
}

func (u *User) Validate() bool {
	if StrongPassword(u.Password) && ValidEmail(u.Email) && Unique(u.Name) {
		return true
	}
	return false
}

func Unique(name string) bool {
	if _, exists := d.Accounts[name]; exists {
		return false
	}
	return true
}

func StrongPassword(password string) bool {
	n := len(password)
	hasLower := false
	hasUpper := false
	hasDigit := false
	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		}
		if ch >= 'a' && ch <= 'z' {
			hasLower = true
		}
		if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}
	if hasLower && hasDigit && hasUpper && (n >= 8) {
		return true
	}
	return false
}

func ValidEmail(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	return false
}

func HashPassword(password string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	encoded := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encoded
}

func Authenticate(u User) bool {
	hashed := HashPassword(u.Password, u.Email)
	if hash, exists := d.Accounts[u.Name]; exists {
		if hash == hashed {
			return true
		}
	}
	return false
}
