package main

import (
	"fmt"
	"learn-crud/api/login"
	"learn-crud/api/register"
	"learn-crud/api/users"
	auth "learn-crud/pkg/middleware/authentication"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", register.SignUp)
	mux.HandleFunc("/login", login.Login)
	mux.Handle("/users", auth.IsAuthenticated(http.HandlerFunc(users.Users)))
	fmt.Println("Server run at http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}
