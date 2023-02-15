package users

import (
	"encoding/json"
	"fmt"
	"learn-crud/pkg/structs"
	"net/http"
)

func Users(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application/json")
		ctx := req.Context()
		decoded := ctx.Value("decoded").(structs.DecodedJWT)
		fmt.Println("The Decoded : ", decoded)

		fmt.Println("You're authenticated..")

		var DB = structs.DB

		result, _ := json.Marshal(DB)
		res.Write(result)
		return
	} else {
		http.Error(res, "This's should be GET request..", 400)
		return
	}
}
