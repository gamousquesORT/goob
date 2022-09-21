package domain

import("errors")

type UserData struct {
	Name string
	Email string
	Password string
	IsAdmin bool
}

const (
	MaxValidName = 20
	MinValidName = 10

	MaxValidPassword = 30
	MinValidPassword = 10
)

var ErrExistingEmail = errors.New("email already exists")
var ErrInvalidUserName = errors.New("user name must have more than 10 and less than 20 chars")
var ErrInvalidUserPassword = errors.New("user password should have more than 10 and less than 30 chars")
var ErrUnknownUserError  = errors.New("unknown user error")


var userEmails = map[string]UserData{}


func NewUser(name string, email string, password string) (*UserData, error) {
	if !checkValidaName(name) {
		return &UserData{}, ErrInvalidUserName
	} else if checkExistingEmail(email) {
		return &UserData{}, ErrExistingEmail
	} else if !checkValidPassword(password) {
		return &UserData{}, ErrInvalidUserPassword
	} 
	ud := new (UserData)
	ud.Name = name
	ud.Email = email
	ud.Password = password
	ud.IsAdmin = false
	userEmails[email] = *ud

	return ud, nil

}

func checkValidaName(name string) bool {
	return len(name) >= MinValidName && len(name) <= MaxValidName
}

func checkValidPassword(password string) bool {
	return len(password) >= MinValidPassword && len(password) <= MaxValidPassword
}

func checkExistingEmail(email string) bool {
	_, ok := userEmails[email]
	return ok
}

func (u *UserData) SetAdmin(admin bool) {
	u.IsAdmin = admin
}

func (u UserData) GetAmdin() bool {
	return u.IsAdmin
}

