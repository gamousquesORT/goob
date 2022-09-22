// TODO 
//  - Too many error types needs refactor
// 
package domain

import (
	"errors"
	"regexp"
)

type UserData struct {
	Name string
	Email string
	Password string
	IsAdmin bool
	Profiles []ProfileData
}

const (
	MaxValidName = 20
	MinValidName = 10

	MaxValidPassword = 30
	MinValidPassword = 10
)

var ErrExistingUserEmail = errors.New("email already exists")
var ErrInvalidUserEmail = errors.New("invalid email format")
var ErrInvalidUserName = errors.New("user name must have more than 10 and less than 20 chars")
var ErrInvalidUserPassword = errors.New("user password should have more than 10 and less than 30 chars")
var ErrInvalidUserError  = errors.New("invalid user error")
var ErrMorethanOneOwner = errors.New("a user can only have one owner profile")
var ErrUnknownUserError  = errors.New("unknown user error")
var ErrTooManyProfiles = errors.New("too many profiles for a user")
var ErrInvalidProfileSequence = errors.New("first profiel should be owner")


var userEmails = map[string]UserData{}


func NewUser(name string, email string, password string) (*UserData, error) {
	if !checkValidaName(name) {
		return &UserData{}, ErrInvalidUserName
	} else if checkExistingEmail(email) {
		return &UserData{}, ErrExistingUserEmail
	} else if !checkInvalidUserEmail(email) {
		return &UserData{}, ErrInvalidUserEmail
	} else if !checkValidPassword(password) {
		return &UserData{}, ErrInvalidUserPassword
	} 
	ud := new (UserData)
	ud.Name = name
	ud.Email = email
	ud.Password = password
	ud.IsAdmin = false
	userEmails[email] = *ud
	//ud.Profiles =  make([]ProfileData, 4)
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

var validEmaiRegExp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)  //^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$
func checkInvalidUserEmail(email string) bool {
	matched := validEmaiRegExp.MatchString(email)

	return matched
}

func (u *UserData) SetAdmin(admin bool) {
	u.IsAdmin = admin
}

func (u UserData) GetAmdin() bool {
	return u.IsAdmin
}

func (u *UserData) AddProfile(alias string, pin int, own bool) error {
	profile, err := NewProfile(alias, pin, own)
	userProfiles := len(u.Profiles)

	if err != nil {
		return err
	} else if userProfiles == 0 && !own {
		return ErrInvalidProfileSequence
	} else if userProfiles == 1 {
		if u.Profiles[0].IsOwnerProfile() && own {
			return ErrMorethanOneOwner
		}
	} else if len(u.Profiles) == 4 {
		return ErrTooManyProfiles
	} 
	u.Profiles = append(u.Profiles, *profile)
	return nil
}

func (u UserData) GetProfile(index int) ProfileData {
	return u.Profiles[index]
}


func (u UserData) GetProfiles() ([]ProfileData) {
	return u.Profiles
}
