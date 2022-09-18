package domain

import("errors")

type UserData struct {
	Name string
	Email string
	Password string
	IsAdmin bool
}

type User interface {
	NewUser(name, email, password string) (UserData, error)
}

var ErrExistingEmail = errors.New("email already exists")
var ErrInvalidUserName = errors.New("user name must be at least 10 characters long")

var userEmails = map[string]UserData{}

func NewUser(name string, email string, password string) (*UserData, error) {
	if !checkExistingEmail(email) && len(name) >= 10 && len(name) <= 20 {
		ud := new (UserData)
		ud.Name = name
		ud.Email = email
		ud.Password = password
		ud.IsAdmin = false
		userEmails[email] = *ud

		return ud, nil

	} else if len(name) < 10 || len(name) > 20 { 
		return &UserData{}, ErrInvalidUserName
	} else {	
		return &UserData{}, ErrExistingEmail
	}

}

func checkExistingEmail(email string) bool {
	_, ok := userEmails[email]
	return ok
}

