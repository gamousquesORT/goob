//
// TODO add poster

package domain

import "errors"

type Rating int

const (  
	G Rating = iota  
	PG Rating = iota  
)


type MainGenre struct {
	gener GenresData
	mainGenres bool
}

type FilmData struct {
	Name string
	genres []MainGenre
	descr string
	rate Rating
	sponsored bool
	// add poster Image
}


var ErrFilmNameMissing  = errors.New("missing film name")

//domain.NewFilmData("Matrix", interface{}, interface{}, "pelicula sci-fi buena parte 1", domain.G, true)
func NewFilmData(name string, gen GenresData, poster any, desc string, r Rating, spon bool) (*FilmData, error) {
	if len(name) < 1 {
		return &FilmData{}, ErrFilmNameMissing
	}

	fd := new(FilmData)
	return fd, nil
}