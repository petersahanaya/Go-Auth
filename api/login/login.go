package login

import (
	"encoding/json"
	"learn-crud/pkg/find"
	bcrypt "learn-crud/pkg/hash"
	"learn-crud/pkg/jwt"
	"learn-crud/pkg/structs"
	"net/http"
	"time"
)

func Login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		res.Header().Set("Content-Type", "application/json")
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

		decodedUser := structs.DecodedJWT{
			Id:       float64(result.Id),
			Username: result.Username,
		}

		ACCESS_TOKEN, _ := jwt.Sign(&decodedUser, time.Now().Add(time.Minute*15).Unix(), "ACCESS_SECRET")
		REFRESH_TOKEN, _ := jwt.Sign(&decodedUser, time.Now().Add(time.Hour*24*7).Unix(), "REFRESH_SECRET")

		http.SetCookie(res, &http.Cookie{
			Name:     "ACCESS_TOKEN",
			Value:    ACCESS_TOKEN,
			Expires:  time.Now().Add(time.Minute * 15),
			HttpOnly: true,
		})

		http.SetCookie(res, &http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    REFRESH_TOKEN,
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			HttpOnly: true,
		})

		res.Write([]byte(`{"msg" : "successfuly login.."}`))

		return
	} else {
		http.Error(res, "This's should be POST request..", 400)
		return
	}
}
