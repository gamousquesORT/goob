//
// TODO add poster

package domain

import "errors"

type Rating int

const (  
	G Rating = iota  
	PG Rating = iota  
)

type Gentype int
const (  
	MainGenre Gentype = iota  
	SecondaryGenre Gentype = iota  
)


type FilmGenre struct {
	Gener GenresData
	GnType Gentype 
}

type FilmData struct {
	Name string
	Descr string
	Rate Rating
	Sponsored bool
	Genres []FilmGenre
	// add poster Image
}


var ErrFilmNameMissing  = errors.New("missing film name")

//domain.NewFilmData("Matrix", interface{}, interface{}, "pelicula sci-fi buena parte 1", domain.G, true)
func NewFilmData(name string, gen GenresData, gentype Gentype, poster string, desc string, r Rating, spon bool) (*FilmData, error) {
	if len(name) < 1 {
		return &FilmData{}, ErrFilmNameMissing
	}

	fd := new(FilmData)
	fd.Name = name
	fd.Descr = desc
	fg := FilmGenre{gen,gentype}
	fd.Genres = make([]FilmGenre,4)
	fd.Genres[0] = fg
	fd.Rate = r
	fd.Sponsored = spon

	return fd, nil
}