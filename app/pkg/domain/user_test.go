package domain_test

import (
	"goob/app/pkg/domain"
	"reflect"
	"testing"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestValidUser(t *testing.T) {

	t.Run("Should return no error given a valid User", func(t *testing.T) {

		got, err := domain.NewUser("uniquename123", "unique@email.com", "uniquepassword")

		prof := []*domain.ProfileData{}
		want := domain.UserData{"uniquename123", "unique@email.com", "uniquepassword", false, prof}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

	t.Run("Should return an error given an exitent email", func(t *testing.T) {
		domain.NewUser("uniquename123", "gaston1@example.com", "uniquepassword")
		_, err := domain.NewUser("uniquename222", "gaston1@example.com", "12345678901")

		assertError(t, err, domain.ErrExistingUserEmail)

	})

	t.Run("Should return an error given an invalid email missing .com", func(t *testing.T) {

		_, err := domain.NewUser("uniquename222", "gaston11@example", "12345678901")

		assertError(t, err, domain.ErrInvalidUserEmail)

	})

	t.Run("Should return an error given an invalid email missing @ ", func(t *testing.T) {

		_, err := domain.NewUser("uniquename222", "gaston11example.com", "12345678901")

		assertError(t, err, domain.ErrInvalidUserEmail)

	})

	t.Run("Should return an error given a user name with less than 10 chars", func(t *testing.T) {
		_, err := domain.NewUser("gaston", "gaston2@example.com", "password")

		assertError(t, err, domain.ErrInvalidUserName)

	})

	t.Run("Should return an error given a user name with less more than 20 chars", func(t *testing.T) {
		_, err := domain.NewUser("123456789012345678901", "gaston3@example.com", "password")

		assertError(t, err, domain.ErrInvalidUserName)

	})

	t.Run("Should return an error given a user password with less than 10 chars", func(t *testing.T) {
		_, err := domain.NewUser("12345678901", "gaston@example4.com", "1")

		assertError(t, err, domain.ErrInvalidUserPassword)

	})

	t.Run("Should return an error given a user name with less more than 20 chars", func(t *testing.T) {
		_, err := domain.NewUser("12345678901", "gaston5@example.com", "1234567890123456789012345678901")

		assertError(t, err, domain.ErrInvalidUserPassword)

	})

}

func TestUserInteractioWithProfile(t *testing.T) {

	t.Run("Should return no error adding the first Profile for a User", func(t *testing.T) {

		user, _ := domain.NewUser("uniquename123", "unique@email.com", "uniquepassword")
		err := user.AddProfile("Alias1", 12345, true)

		assertNoError(t, err)
	})

	t.Run("Should return no error adding the first Profile for a User and getting it back", func(t *testing.T) {

		user, _ := domain.NewUser("uniquename123", "unique@email.com", "uniquepassword")
		want, _ := domain.NewProfile("Alias1", 12345, true)
		user.AddProfile(want.Alias, want.Pin, want.Owner)
		got := user.GetProfile(0)

		if reflect.DeepEqual(*got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

	})

	t.Run("Should return an error setting more than one owner profile", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2345@email.com", "uniquepassword")

		user.AddProfile("Alias1", 12345, true)
		err := user.AddProfile("Alias2", 12345, true)
		assertError(t, err, domain.ErrMoreThanOneOwner)
	})

	t.Run("Should return an error if first profile is not owner", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2355@email.com", "uniquepassword")

		user.AddProfile("Alias1", 12345, false)
		err := user.AddProfile("Alias2", 12345, false)
		assertError(t, err, domain.ErrInvalidProfileSequence)
	})

	t.Run("Should return an error setting more than four  profiles", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2346@email.com", "uniquepassword")
		user.AddProfile("Alias1", 12345, true)
		user.AddProfile("Alias2", 12345, false)
		user.AddProfile("Alias3", 12345, false)
		user.AddProfile("Alias4", 12345, false)
		err := user.AddProfile("Alias5", 12345, false)
		assertError(t, err, domain.ErrTooManyProfiles)
	})

	t.Run("Should return an list with profiles", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2347@email.com", "uniquepassword")
		user.AddProfile("Alias1", 12345, true)
		user.AddProfile("Alias2", 12345, false)

		want := []*domain.ProfileData{{Alias: "Alias1", Pin: 12345, Owner: true, Child: false, FilmsDetails: map[string]*domain.ProfileFilmDetails{}}, {Alias: "Alias2", Pin: 12345, Owner: false, Child: false, FilmsDetails: map[string]*domain.ProfileFilmDetails{}}}

		got := user.GetProfiles()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}
	})

	t.Run("Should return an error given a duplicated profile alias", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2347@email.com", "uniquepassword")
		user.AddProfile("Alias1", 12345, true)
		err := user.AddProfile("Alias1", 12345, false)
		assertError(t, err, domain.ErrDuplicatedAlias)
	})

	t.Run("Should NOT be able to set a child profile been a not owner profile", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2347@email.com", "uniquepassword")
		err := user.AddProfile("Alias1", 12345, true)

		assertNoError(t, err)

		user.SetChildProfile("Alias1")

		err2 := user.IsChildProfile("Alias21")

		assertError(t, err2, domain.ErrInvalidProfileAction)
	})
}
