package domain_test

import ("testing"
"streamapp.com/domain"
"reflect"
)


func TestValidGenres(t *testing.T) {

	t.Run("Should return no error given a valid Genres data", func(t *testing.T) {
		got, err := domain.NewGenresData("Terror",  "Genero para tener miedo")

		want := domain.GenresData{"Terror",  "Genero para tener miedo"}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})
	

	t.Run("Should return no error given a valid Genres data", func(t *testing.T) {
		_, err := domain.NewGenresData("", "Genero para entretenerse")
	
		assertNoError(t, err)
	
	})
		/*
	t.Run("Should return an error given a Film without name", func(t *testing.T) {
		_, err := domain.NewFilmData("", 1, 1, "pelicula sci-fi buena parte 1", domain.G, true)
	
		assertError(t, err, domain.ErrFilmNameMissing)
	})
	*/
}