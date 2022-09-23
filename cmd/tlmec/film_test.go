package domain_test

import ("testing"
"streamapp.com/domain"
"reflect"
)

func TestValidFilm(t *testing.T) {

	t.Run("Should return no error given a valid Film data", func(t *testing.T) {
		g := domain.GenresData{}
		got, err := domain.NewFilmData("Matrix", g, 1, "pelicula sci-fi buena parte 1", domain.G, true)
		want := domain.FilmData{}
		if !reflect.DeepEqual(*got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})
	
	t.Run("Should return no error given a valid Film data", func(t *testing.T) {
		g := domain.GenresData{}
		got, err := domain.NewFilmData("Matrix", g, 1, "pelicula sci-fi buena parte 1", domain.G, true)
		want := domain.FilmData{}
		if !reflect.DeepEqual(*got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})
	
	t.Run("Should return an error given a Film without name", func(t *testing.T) {
		g := domain.GenresData{}
		_, err := domain.NewFilmData("", g, 1, "pelicula sci-fi buena parte 1", domain.G, true)
	
		assertError(t, err, domain.ErrFilmNameMissing)
	})
}