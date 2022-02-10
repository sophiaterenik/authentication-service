package handler

import (
	"authentication-service/pkg/account"
	"crypto/sha256"
	_ "crypto/sha256"
	"encoding/base64"
	_ "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from home")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// collect username, email address, and password
	// 1. check password meets complexity requirements
	// 2. OPTIONAL use regular expression to check email
	// 3. check username is unique
	// 4. hash the password with email address as salt
	// 5. store the account without the email address
	var u account.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username: [%v], email: [%v]", u.Name, u.Email)

	// 1. check password
	if !StrongPassword(u.Password) {
		//throw error
	}

	// 2. check email
	if !CheckEmail(u.Email) {
		//throw error
	}

	// 4. hash password
	hashedPassword := HashPassword(u.Password, u.Email)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 1. collect username, email, and password from r
	// 2. check if arguments are empty
	// 3. hash the provided email and password using same algorithm
	// 4. lookup hash code by provided username
	// 5. check if stored hash code matches the generated hash code
	var u account.User
	// 1. collect arguments
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from login")

	// 2. check if empty
	if u.Email == "" {
		//throw error
	}
	if u.Name == "" {
		//throw error
	}
	if u.Password == "" {
		//throw error
	}
	hashedPassword := HashPassword(u.Password, u.Email)
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

func CheckEmail(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	return false
}

func HashPassword(password string, salt string) string {
	passwordBytes := []byte(password)
	saltBytes := []byte(password)
	sha256Hasher := sha256.New()
	passwordBytes = append(passwordBytes, saltBytes...)
	sha256Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha256Hasher.Sum(nil)
	encodedPassword := base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return encodedPassword
}
