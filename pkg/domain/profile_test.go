package domain_test

import (
	domain "goob/domain/pkg/domain"
	"reflect"
	"testing"
)

func TestValidOwnerProfile(t *testing.T) {

	t.Run("Should return no error creating new valid profile", func(t *testing.T) {
		got, err := domain.NewProfile("alias", 12345, true)

		films := []domain.ProfileFilmDetails{}
		want := domain.ProfileData{"alias", 12345, true, false, films}

		if !reflect.DeepEqual(got, &want) {
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

	t.Run("Should return no error creating an owner profile with true flag", func(t *testing.T) {

		g, err := domain.NewProfile("1234567890", 12345, true)
		got := g.Owner
		want := true

		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

	t.Run("Should be able to set a child profile been an owner", func(t *testing.T) {

		g, err := domain.NewProfile("1234567890", 12345, true)

		g.SetChildProfile(true)
		got := g.IsChildProfile()
		want := true

		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

}

func TestProfileFilmInteraction(t *testing.T) {

	t.Run("Should return no error adding a film", func(t *testing.T) {

		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		err := prof.AddFilm(film)

		assertNoError(t, err)
	})

	t.Run("Should return Added Film given a film", func(t *testing.T) {

		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		err := prof.AddFilm(film)

		got := prof.GetFilmsDetails()

		want := []domain.ProfileFilmDetails{{Film: film, Votes: 0}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

}
