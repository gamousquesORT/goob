package domain_test

import (
	"goob/app/pkg/domain"
	"reflect"
	"testing"
)



func TestAdminServices(t *testing.T) {
	t.Run("Should create the default Admin User ", func(t *testing.T) {
		var streamApp = domain.NewStreamApp()
		admin, err := streamApp.CreateAdmin()
		got := admin.IsAdmin
		want := true 
		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}
		assertNoError(t, err)
		
	})

	t.Run("Should create a new genre with no error given valid data ", func(t *testing.T) {
		var streamApp = domain.NewStreamApp()
		 err := streamApp.CreateGenre("Terror", "Genero para tener miedo")

		assertNoError(t, err)
	})

	t.Run("Should create a new genre and retrieve it ", func(t *testing.T) {
		var streamApp = domain.NewStreamApp()
		err := streamApp.CreateGenre("Terror", "Genero para tener miedo")
		got := streamApp.GetGenres()
		want := []domain.GenresData{{Name: "Terror", Description: "Genero para tener miedo"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)}

	   assertNoError(t, err)
   })


}
