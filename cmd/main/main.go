package main

import (
	"fmt"
	"learn-crud/api/login"
	"learn-crud/api/register"
	"net/http"
)

func main() {
	http.HandleFunc("/register", register.SignUp)
	http.HandleFunc("/login", login.Login)
	fmt.Println("Server run at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
