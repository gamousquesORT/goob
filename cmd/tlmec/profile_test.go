package domain_test


import (
	"testing"
	"streamapp.com/domain"
)


func TestValidProfile(t *testing.T) {

	t.Run("Should return no error creating new vald profile", func(t *testing.T) {

		got, err := domain.NewProfile("alias", 12345)

		want := domain.ProfileData{"alias", 12345}

		if *got  != want {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})



}