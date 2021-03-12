package auth

type Auth struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func NewAuth(
	ID int,
	name string,
	email string,
	password string,
) Auth {
	return Auth{
		ID:       ID,
		Name:     name,
		Email:    email,
		Password: password,
	}
}
