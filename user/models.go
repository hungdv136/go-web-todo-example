package user

type User struct {
	Id           int
	Username     string
	Name         string
	PasswordHash string
}

type LoginModel struct {
	Username string
	Password string
}
