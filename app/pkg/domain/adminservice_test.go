package domain_test

import (
	"goob/app/pkg/domain"
	"reflect"
	"testing"
)

func TestAdminServices(t *testing.T) {
	t.Run("Should create a new genre with no error given valid data ", func(t *testing.T) {
		var streamApp = domain.NewStreamApp()
		err := streamApp.CreateGenre("Terror", "Genero para tener miedo")

		assertNoError(t, err)
	})

	t.Run("Should create a new genre and retrieve it ", func(t *testing.T) {
		var streamApp = domain.NewStreamApp()
		streamApp.CreateGenre("Terror", "Genero para tener miedo")
		got, err := streamApp.GetGenres("Terror")
		want := domain.GenresData{Name: "Terror", Description: "Genero para tener miedo"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

}
