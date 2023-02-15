package structs

var DB []User

type User struct {
	Id       int
	Username string
	Password string
}

type DecodedJWT struct {
	Id       float64
	Username string
	Expires  float64
}
