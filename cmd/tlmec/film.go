package domain

type Rating int

const (  
	G Rating = iota  
	PG Rating = iota  
)

type FilmData struct {

}

//domain.NewFilmData("Matrix", interface{}, interface{}, "pelicula sci-fi buena parte 1", domain.G, true)
func NewFilmData(name string, genres any, poster any, desc string, r Rating, spon bool) (*FilmData, error) {
	fd := new(FilmData)
	return fd, nil
}