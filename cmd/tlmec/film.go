package domain

import "errors"

type Rating int

const (  
	G Rating = iota  
	PG Rating = iota  
)

type FilmData struct {

}


var ErrFilmNameMissing  = errors.New("missing film name")

//domain.NewFilmData("Matrix", interface{}, interface{}, "pelicula sci-fi buena parte 1", domain.G, true)
func NewFilmData(name string, genres any, poster any, desc string, r Rating, spon bool) (*FilmData, error) {
	if len(name) < 1 {
		return &FilmData{}, ErrFilmNameMissing
	}

	fd := new(FilmData)
	return fd, nil
}