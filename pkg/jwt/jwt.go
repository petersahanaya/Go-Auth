package jwt

import (
	"fmt"
	"learn-crud/pkg/env"
	"learn-crud/pkg/structs"

	jwt "github.com/golang-jwt/jwt/v4"
)

func Sign(user *structs.DecodedJWT, long int64, types string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       user.Id,
		"Username": user.Username,
		"exp":      long,
	})

	if types == "ACCESS_SECRET" {
		tokenString, err := token.SignedString([]byte(env.GetEnv("ACCESS_SECRET")))

		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}

		return tokenString, nil
	} else {
		tokenString, err := token.SignedString([]byte(env.GetEnv("REFRESH_SECRET")))

		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}

		return tokenString, nil
	}
}

func Verify(value string, types string) (structs.DecodedJWT, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("%v", "You're not Authorized..")
		}

		if types == "ACCESS_SECRET" {
			return []byte(env.GetEnv("ACCESS_SECRET")), nil
		} else {
			return []byte(env.GetEnv("REFRESH_SECRET")), nil
		}
	})

	if err != nil {
		return structs.DecodedJWT{}, fmt.Errorf("%v", "You're not Authorized..")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return structs.DecodedJWT{}, fmt.Errorf("%v", "You're not Authorized..")
	}

	user := structs.DecodedJWT{
		Id:       claims["Id"].(float64),
		Username: claims["Username"].(string),
		Expires:  claims["exp"].(float64),
	}

	return user, nil
}
