package handler

import (
	"authentication-service/pkg/account"
	_ "crypto/sha256"
	_ "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// home page

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from home")
}

// sign up page

func SignUp(w http.ResponseWriter, r *http.Request) {

	var u account.User
	// check if user is null
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// check if user is valid
	if !u.Validate() {
		http.Error(w, "Account is not valid", http.StatusBadRequest)
		return
	}
	// hash password with email as salt
	hashed := account.HashPassword(u.Password, u.Email)
	// store user in map
	account.Store(u.Name, hashed)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username: [%v], email: [%v]", u.Name, u.Email)
}

// login page

func Login(w http.ResponseWriter, r *http.Request) {

	var u account.User
	// check if user is null
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// check if user is in map
	if !account.Authenticate(u) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from login")
}
