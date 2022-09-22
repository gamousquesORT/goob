package domain_test

import ("testing"
"streamapp.com/domain"
"reflect"
)

func TestValidFilm(t *testing.T) {

	t.Run("Should return no error given a valid Film data", func(t *testing.T) {
		got, err := domain.NewFilmData("Matrix", 1, 1, "pelicula sci-fi buena parte 1", domain.G, true)
		want := domain.FilmData{}
		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})
}