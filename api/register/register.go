package register

import (
	"encoding/json"
	"fmt"
	"learn-crud/pkg/find"
	bcrypt "learn-crud/pkg/hash"
	"learn-crud/pkg/structs"
	"net/http"
)

func SignUp(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		res.Header().Set("Content-Type", "application/json")
		var user structs.User

		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

		fmt.Println(user)

		duplicate, _ := find.FindOne(user.Username, &structs.DB)

		if duplicate {
			http.Error(res, "There's another user used this username..", 400)
		}

		hashPass, _ := bcrypt.HashPassword(&user.Password)

		structs.DB = append(structs.DB, structs.User{Id: len(structs.DB) + 1, Username: user.Username, Password: hashPass})

		fmt.Fprint(res, `{"msg" : "User Register Success"}`)
		fmt.Println(structs.DB)
		return
	} else {
		http.Error(res, "This's should be POST request..", 400)
		return
	}
}
