package find

import (
	"learn-crud/pkg/structs"
	"strings"
)

func FindOne(username string, DB *[]structs.User) (bool, structs.User) {
	var result bool
	var users structs.User
	for _, user := range *DB {
		if strings.EqualFold(username, user.Username) {
			result = true
			users = user
			break
		}
	}

	return result, users
}
