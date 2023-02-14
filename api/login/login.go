package login

import (
	"encoding/json"
	"fmt"
	"learn-crud/pkg/find"
	bcrypt "learn-crud/pkg/hash"
	"learn-crud/pkg/structs"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var user structs.User

		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		finded, result := find.FindOne(user.Username, &structs.DB)

		if !finded {
			http.Error(res, "Cannot find user..", 400)
			return
		}

		authenticated := bcrypt.ComparePassword(&result.Password, &user.Password)

		if !authenticated {
			http.Error(res, "Wrong Password..", 400)
			return
		}

		fmt.Fprintf(res, "%v successfuly login", result.Username)
		return
	} else {
		http.Error(res, "This's should be POST request..", 400)
		return
	}
}
