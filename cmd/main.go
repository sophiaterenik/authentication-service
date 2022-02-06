package main

import (
	"authentication-service/pkg/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/signup", handler.SignUp)
	r.HandleFunc("/login", handler.Login)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
