//
// TODO add poster

package domain

import (
	"errors"
	"strconv"
)

type Rating int

const (
	G  Rating = iota
	PG Rating = iota
)

type Gentype int

const (
	MainGenre      Gentype = iota
	SecondaryGenre Gentype = iota
)

const MaxNumberOfGenres int = 4

type FilmGenre struct {
	Genre  GenresData
	GnType Gentype
}

type FilmData struct {
	Name        string
	Description string
	Rate        Rating
	Sponsored   bool
	Genres      []FilmGenre
	// add poster Image
}

var ErrFilmNameMissing = errors.New("missing film name")
var ErrExpectedPrimaryGenre = errors.New("missing primary genre")
var ErrTooManyGenres = errors.New("max number of genres is " + strconv.Itoa(MaxNumberOfGenres))

// domain.NewFilmData("Matrix", interface{}, interface{}, "pelicula sci-fi buena parte 1", domain.G, true)
func NewFilmData(name string, gen GenresData, gentype Gentype, poster string, desc string, r Rating, spon bool) (*FilmData, error) {
	if len(name) < 1 {
		return &FilmData{}, ErrFilmNameMissing
	} else if gentype != MainGenre {
		return &FilmData{}, ErrExpectedPrimaryGenre
	}

	fd := new(FilmData)
	fd.Name = name
	fd.Description = desc
	fg := FilmGenre{gen, gentype}
	fd.Genres = make([]FilmGenre, 1)
	fd.Genres[0] = fg
	fd.Rate = r
	fd.Sponsored = spon

	return fd, nil
}

func (film *FilmData) AddGenre(genre GenresData, genreType Gentype) error {
	if len(film.Genres) == MaxNumberOfGenres {
		return ErrTooManyGenres
	}
	filmGen := FilmGenre{}
	filmGen.Genre = genre
	filmGen.GnType = genreType

	destSlice := append(film.Genres, filmGen)
	film.Genres = destSlice

	return nil
}

func (film *FilmData) GetGenres() []FilmGenre {
	return film.Genres
}
