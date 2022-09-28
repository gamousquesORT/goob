package domain

import "errors"


type GenresData struct {
	Name string
	Description string

}

var ErrInvalidGenresData = errors.New("invalid Genres name")

func NewGenresData(name string,  desc string) (*GenresData, error) {
	if len(name) == 0 {
		return &GenresData{}, ErrInvalidGenresData
	}
	
	gd := new(GenresData)
	gd.Name = name
	gd.Description = desc

	return gd, nil
}