package handler

import (
	"authentication-service/pkg/account"
	_ "crypto/sha256"
	_ "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
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
	if !u.Validate() {
		http.Error(w, "Account is not valid", http.StatusBadRequest)
		return
	}
	hashed := account.HashPassword(u.Password, u.Email)
	account.Store(u.Name, hashed)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username: [%v], email: [%v]", u.Name, u.Email)
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

	if !account.Authenticate(u) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from login")
}
