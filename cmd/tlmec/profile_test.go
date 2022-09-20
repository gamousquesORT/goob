package domain_test


import (
	"testing"
	"streamapp.com/domain"
)


func TestValidOwnerProfile(t *testing.T) {

	t.Run("Should return no error creating new valid profile", func(t *testing.T) {

		got, err := domain.NewProfile("alias", 12345, true)

		want := domain.ProfileData{"alias", 12345, true}

		if *got  != want {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})


	t.Run("Should return an error creating new profile with short Alias", func(t *testing.T) {

		_, err := domain.NewProfile("", 12345, true)

		assertError(t, err, domain.ErrEInvalidAlias)
	})

	t.Run("Should return an error creating new profile with long Alias", func(t *testing.T) {

		_, err := domain.NewProfile("1234567890123456", 12345, true)

		assertError(t, err, domain.ErrEInvalidAlias)
	})

	t.Run("Should return an error creating new profile with short Pin", func(t *testing.T) {

		_, err := domain.NewProfile("1234567890", 1234, true)

		assertError(t, err, domain.ErrEInvalidPin)
	})

	t.Run("Should return an error creating new profile with long Pin", func(t *testing.T) {

		_, err := domain.NewProfile("1234567890", 123456, true)

		assertError(t, err, domain.ErrEInvalidPin)
	})

	t.Run("Should return an error creating new owner profile with false flag", func(t *testing.T) {

		g, err := domain.NewProfile("1234567890", 12345, true)
		got := g.Owner
		want := true

		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

}