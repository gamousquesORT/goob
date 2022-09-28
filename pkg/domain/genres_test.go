package domain_test

import (
	"goob/app/pkg/domain"
	"reflect"
	"testing"
)

func TestValidGenres(t *testing.T) {

	t.Run("Should return no error given a valid Genres data", func(t *testing.T) {
		got, err := domain.NewGenresData("Terror", "Genero para tener miedo")

		want := domain.GenresData{"Terror", "Genero para tener miedo"}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)

	})

	t.Run("Should return  error given an empty Genres description", func(t *testing.T) {
		_, err := domain.NewGenresData("", "Genero para entretenerse")

		assertError(t, err, domain.ErrInvalidGenresData)

	})

}
