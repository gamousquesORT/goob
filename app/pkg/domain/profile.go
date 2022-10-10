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
	FilmsDetails map[string]*ProfileFilmDetails
}

type ProfileFilmDetails struct {
	Film *FilmData
	Vote UserRating
	Watched bool
}

const (
	MinValidAlias = 1
	MaxValidAlias = 15
)

type  UserRating int

const (
	ThumbDown UserRating = iota
	Thumbup UserRating = iota
	DoubleThumbsUp UserRating = iota
)

// TODO chceck profiles errors, too many
var ErrEInvalidAlias = errors.New("alias should be greater than 1 and less 16")
var ErrProfilenotFound = errors.New("alias does not exist")
var ErrEInvalidPin = errors.New("alias should be of 5 digits")
var ErrAddingFilmToProfile = errors.New("could'nt add film to profile")
var ErrInvalidProfileAction = errors.New("invalid action with owner profle")
var ErrInvalidFilm = errors.New("film does not exist")


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
	val.FilmsDetails = make(map[string]*ProfileFilmDetails)

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
	fd := ProfileFilmDetails{&film, 0, false}
	p.FilmsDetails[film.Name] = &fd
	return nil
}

func (p ProfileData) GetFilmsDetails(film FilmData) (ProfileFilmDetails, error) {
	fd, ok := p.FilmsDetails[film.Name]
	if !ok {
		return ProfileFilmDetails{}, ErrInvalidFilm
	}
	return  *fd, nil
}

func (p *ProfileData) RateFilm(film FilmData, rating UserRating) error {
	fd, ok := p.FilmsDetails[film.Name]
	if !ok {
		return ErrInvalidFilm
	}
	
	fd.Vote = rating
	
	return nil

}

func (p ProfileData) GetFilmUserRating(film FilmData) (UserRating, error) {
	fd, ok := p.FilmsDetails[film.Name]
	if !ok {
		return -1, ErrInvalidFilm
	}

	return fd.Vote, nil
}

func (p* ProfileData) MarkAsWatched(film FilmData) error {
	fd, ok := p.FilmsDetails[film.Name]
	if !ok {
		return ErrInvalidFilm
	}
	fd.Watched = true
	return nil
}

func (p ProfileData) Watched(film FilmData) (bool, error) {
	fd, ok := p.FilmsDetails[film.Name]
	if !ok {
		return false, ErrInvalidFilm
	}
	return fd.Watched, nil
}