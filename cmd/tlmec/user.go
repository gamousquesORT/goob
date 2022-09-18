package domain

import("errors")

type UserData struct {
	Name string
	Email string
	Password string
	IsAdmin bool
}


var ErrExistingEmail = errors.New("email already exists")
var ErrInvalidUserName = errors.New("user name must have more than 10 and less than 20 chars")
var ErrInvalidUserPassword = errors.New("user password should have more than 10 and less than 30 chars")

var userEmails = map[string]UserData{}

func NewUser(name string, email string, password string) (*UserData, error) {
	if !checkExistingEmail(email) && chackValidaName(name) && checkValidPassword(password) {
		ud := new (UserData)
		ud.Name = name
		ud.Email = email
		ud.Password = password
		ud.IsAdmin = false
		userEmails[email] = *ud

		return ud, nil

	} else if len(name) < 10 || len(name) > 20 { 
		return &UserData{}, ErrInvalidUserName
	} else if (len(password) < 10 || len(password) > 30) {
		return &UserData{}, ErrInvalidUserPassword
	} else {	
		return &UserData{}, ErrExistingEmail
	}

}

func chackValidaName(name string) bool {
	return len(name) >= 10 && len(name) <= 20
}

func checkValidPassword(password string) bool {
	return len(password) >= 10 && len(password) <= 30
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

