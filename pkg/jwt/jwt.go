package jwt

import (
	"fmt"
	"learn-crud/pkg/env"
	"learn-crud/pkg/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Sign(user *structs.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       user.Id,
		"Username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString(env.GetEnv("SECRET"))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func Verify(token string) {

}
