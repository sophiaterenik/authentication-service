package account

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

var d = Database{}

// user information

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// store user information

type Database struct {
	Accounts map[string]string
}

// create a map

func (s *Database) Initialize() {
	s.Accounts = make(map[string]string)
}

// store a user in the map

func Store(u string, hash string) {
	if d.Accounts == nil {
		d.Initialize()
	}
	d.Accounts[u] = hash
}

// check if user is valid

func (u *User) Validate() bool {
	if StrongPassword(u.Password) && ValidEmail(u.Email) && Unique(u.Name) {
		return true
	}
	return false
}

// check if username is not already stored

func Unique(name string) bool {
	if _, exists := d.Accounts[name]; exists {
		return false
	}
	return true
}

// check if strong password

func StrongPassword(password string) bool {
	n := len(password)
	hasLower := false
	hasUpper := false
	hasDigit := false
	for _, ch := range password {
		// password has an upper case letter
		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		}
		// password has a lower case letter
		if ch >= 'a' && ch <= 'z' {
			hasLower = true
		}
		// password has a digit
		if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}
	if hasLower && hasDigit && hasUpper && (n >= 8) {
		return true
	}
	return false
}

// check if email is valid

func ValidEmail(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	return false
}

// hash the password with a salt

func HashPassword(password string, salt string) string {
	// create new hasher
	hasher := sha256.New()
	// hash the password with the salt
	hasher.Write([]byte(password + salt))
	// store hashed password
	encoded := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encoded
}

// check if user is in map

func Authenticate(u User) bool {
	// hash the user
	hashed := HashPassword(u.Password, u.Email)
	// search map for user, return true if in map
	if hash, exists := d.Accounts[u.Name]; exists {
		if hash == hashed {
			return true
		}
	}
	return false
}
