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
	FilmsDetails []ProfileFilmDetails
}

type ProfileFilmDetails struct {
	Film FilmData
	Votes int
}

const (
	MinValidAlias = 1
	MaxValidAlias = 15
)

var ErrEInvalidAlias = errors.New("alias should be greater than 1 and less 16")
var ErrEInvalidPin = errors.New("alias should be of 5 digits")
var ErrAddingFilmToProfile = errors.New("could'nt add film to profile")
var ErrInvalidProfileAction = errors.New("invalid action with owner profle")

func checkValidAlias(alias string) bool {
	return len(alias) >= MinValidAlias && len(alias) <= MaxValidAlias
}

func checkValidPin(pin int) bool {
	s := strconv.Itoa(pin)
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	if !b || len(s) != 5 {
		return false
	}
	return true
	
}

func NewProfile(alias string, pin int, own bool) (*ProfileData, error) {
	val := new(ProfileData)
	if !checkValidAlias(alias) {
		return &ProfileData{}, ErrEInvalidAlias
	} else if !checkValidPin(pin) {
		return &ProfileData{},ErrEInvalidPin
	}

	val.Alias = alias
	val.Pin = pin
	val.Owner = own
	val.Child = false
	val.FilmsDetails = []ProfileFilmDetails{}

	return  val,nil
}


func (p *ProfileData) SetChildProfile(value bool)error {
	if p.Owner {
		p.Child = value;
		return nil
	}

	return ErrInvalidProfileAction
}

func (p ProfileData) IsChildProfile() bool {
	return p.Child
}


func (p ProfileData) IsOwnerProfile() bool {
	return p.Owner
}

func (p *ProfileData) AddFilm(film FilmData) error {
	fd := ProfileFilmDetails{film, 0}
	p.FilmsDetails = append(p.FilmsDetails, fd)
	return nil
}

func (p ProfileData) GetFilmsDetails() []ProfileFilmDetails {
	return p.FilmsDetails
}