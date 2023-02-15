package authentication

import (
	"context"
	"fmt"
	"learn-crud/pkg/jwt"
	"net/http"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		token, err := req.Cookie("ACCESS_TOKEN")

		if err != nil {
			res.WriteHeader(401)
			res.Write([]byte("There's something wrong when parse cookie.."))
			return
		}

		decoded, errs := jwt.Verify(token.Value, "ACCESS_SECRET")

		ctx := req.Context()
		fmt.Println("INSIDE MIDDLEWARE : ", decoded)

		ctx = context.WithValue(ctx, "decoded", decoded)

		if errs != nil {
			res.WriteHeader(401)
			res.Write([]byte("You're not Authenticated.."))
			return
		}

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
