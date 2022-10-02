package domain_test

import (
	"goob/app/pkg/domain"
	"reflect"
	"testing"
)

func TestValidOwnerProfile(t *testing.T) {

	t.Run("Should return no error creating new valid profile", func(t *testing.T) {
		got, err := domain.NewProfile("alias", 12345, true)

		films := map[string]*domain.ProfileFilmDetails{}
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

	t.Run("Should NOT be able to set a child profile been a not owner profile", func(t *testing.T) {

		g,_ := domain.NewProfile("1234567890", 12345, false)


		err := g.SetChildProfile(true)
		

		assertError(t, err, domain.ErrInvalidProfileAction)
	})

	t.Run("Should  be able to set a child profile been a not owner profile", func(t *testing.T) {

		g,_ := domain.NewProfile("1234567890", 12345, true)


		err := g.SetChildProfile(true)
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

		prof.AddFilm(film)

		got, err := prof.GetFilmsDetails(film)

		want := domain.ProfileFilmDetails{Film: &film, Vote: domain.ThumbDown}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})


	t.Run("Should return no error given Thumbup for an existing Film", func(t *testing.T) {
		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		prof.AddFilm(film)

		err := prof.RateFilm(film, domain.Thumbup)
		assertNoError(t, err)

	})

	t.Run("Should return an error given Thumb vote for an Non existing Film", func(t *testing.T) {
		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		err := prof.RateFilm(film, domain.Thumbup)
		assertError(t, err, domain.ErrInvalidFilm)

	})

	t.Run("Should return the same Thumb vote for an existing Film", func(t *testing.T) {
		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		prof.AddFilm(film)

		want := domain.Thumbup
		prof.RateFilm(film, want)

		got, err := prof.GetFilmUserRating(film)

		
		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}
		assertNoError(t, err)

	})

	t.Run("Should be able to mark as watch for a Film", func(t *testing.T) {
		prof, _ := domain.NewProfile("alias", 12345, true)
		film := createValidFilm(t)

		prof.AddFilm(film)

		want := true
		prof.MarkAsWatched(film)

		got, err := prof.Watched(film)

		
		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}
		assertNoError(t, err)

	})
}

