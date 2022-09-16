package domain

type UserData struct {
	name string
	email string
	password string
}

type User interface {
	NewUser(name, email, password string) (UserData, error)
}

func NewUser(name string, email string, password string) (*UserData, error) {
	ud := new (UserData)
	ud.name = name
	ud.email = email
	ud.password = password
	
	return ud, nil

}
