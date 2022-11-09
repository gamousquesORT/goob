// TODO
//   - Too many error types needs refactor
package domain

import (
	"errors"
	"regexp"
)

type UserData struct {
	Name     string
	Email    string
	Password string
	IsAdmin  bool
	Profiles []*ProfileData
}

const (
	MaxValidName = 20
	MinValidName = 10

	MaxValidPassword = 30
	MinValidPassword = 10

	MinUserProfiles = 0
	MaxUserProfiles = 4
)

var ErrExistingUserEmail = errors.New("email already exists")
var ErrInvalidUserEmail = errors.New("invalid email format")
var ErrInvalidUserName = errors.New("user name must have more than 10 and less than 20 chars")
var ErrInvalidUserPassword = errors.New("user password should have more than 10 and less than 30 chars")
var ErrInvalidUserError = errors.New("invalid user error")
var ErrMoreThanOneOwner = errors.New("a user can only have one owner profile")
var ErrUnknownUserError = errors.New("unknown user error")
var ErrTooManyProfiles = errors.New("too many profiles for a user")
var ErrInvalidProfileSequence = errors.New("first profiel should be owner")
var ErrDuplicatedAlias = errors.New("duplicated Alias")

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
	ud := new(UserData)
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

var validEmailRegExp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`) //^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$
func checkInvalidUserEmail(email string) bool {
	matched := validEmailRegExp.MatchString(email)

	return matched
}

func (user *UserData) SetAdmin(admin bool) {
	user.IsAdmin = admin
}

func (user *UserData) GetAdmin() bool {
	return user.IsAdmin
}

func (user *UserData) AddProfile(alias string, pin int, own bool) error {
	profile, err := NewProfile(alias, pin, own)

	if err != nil {
		return err
	}

	retVal := validateProfileToAdd(user, *profile)

	if retVal != nil {
		return retVal
	}

	user.Profiles = append(user.Profiles, profile)
	return nil
}

func validateProfileToAdd(user *UserData, profile ProfileData) error {
	userProfiles := len(user.Profiles)
	if checkNoUserProfiles(userProfiles) && !profile.Owner {
		return ErrInvalidProfileSequence
	} else if userProfiles == 1 && user.Profiles[0].IsOwnerProfile() && profile.Owner {
		return ErrMoreThanOneOwner
	} else if checkMaxUserProfiles(userProfiles) {
		return ErrTooManyProfiles
	} else if checkDuplicatedAlias(user, profile.Alias) {
		return ErrDuplicatedAlias
	}
	return nil
}

func checkMaxUserProfiles(userProfiles int) bool {
	return userProfiles == MaxUserProfiles
}

func checkNoUserProfiles(userProfiles int) bool {
	return userProfiles == MinUserProfiles
}

func (user *UserData) GetProfile(index int) *ProfileData {
	return user.Profiles[index]
}

func (user *UserData) GetProfiles() []*ProfileData {
	return user.Profiles
}

func checkDuplicatedAlias(user *UserData, alias string) bool {
	for _, p := range user.Profiles {
		if p.Alias == alias {
			return true
		}
	}
	return false
}

func (user *UserData) SetChildProfile(alias string) error {
	for _, p := range user.Profiles {
		if p.Alias == alias {
			if p.Owner {
				p.SetChildProfile(true)
				return nil
			}

		}
	}

	return ErrInvalidProfileAction

}

func (user *UserData) IsChildProfile(alias string) error {
	for _, p := range user.Profiles {
		if p.Alias == alias {
			if p.Child {
				return nil
			}
		}
	}

	return ErrInvalidProfileAction
}
