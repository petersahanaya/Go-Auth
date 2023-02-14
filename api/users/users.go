package users

import "net/http"

func Users(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

	} else {
		http.Error(res, "This's should be GET request..", 400)
		return
	}
}
