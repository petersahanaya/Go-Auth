package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	value := os.Getenv(name)

	return value
}
