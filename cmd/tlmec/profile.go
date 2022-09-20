package domain

import (
	"errors"

	"strconv"
)

type ProfileData struct {
	Alias string
	Pin int
	Owner bool
	Child bool
}


var ErrEInvalidAlias = errors.New("alias should be greater than 1 and less 16")
var ErrEInvalidPin = errors.New("alias should be of 5 digits")



func NewProfile(alias string, pin int, own bool) (*ProfileData, error) {
	val := new(ProfileData)
	if len(alias) <1 || len(alias) > 15 {
		return &ProfileData{}, ErrEInvalidAlias
	}

	s := strconv.Itoa(pin)
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	if !b || len(s) != 5 {
		return &ProfileData{},ErrEInvalidPin
	}
	
	val.Alias = alias
	val.Pin = pin
	val.Owner = own
	val.Child = false

	return  val,nil
}


func (p *ProfileData) SetChildProfile(value bool) {
	p.Child = value;
}

func (p ProfileData) IsChildProfile() bool {
	return p.Child
}