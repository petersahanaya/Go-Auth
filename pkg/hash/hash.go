package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 10)

	return string(bytes), err
}

func ComparePassword(hashPass, originalPass *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hashPass), []byte(*originalPass))

	return err == nil
}
