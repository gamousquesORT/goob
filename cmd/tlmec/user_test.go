package domain_test

import (
	"reflect"
	"testing"
	"streamapp.com/domain"
)

var adminUser = domain.UserData{}

func createAdminUser(t testing.TB) {
	t.Helper()
	adminUser, _ := domain.NewUser("Admin6789101", "admin@example.com", "password")
	adminUser.SetAdmin(true)

}

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

		prof := []domain.ProfileData{}
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
		prof, _ := domain.NewProfile("Alias1", 12345, true)
		err := user.AddProfile(*prof)

		assertNoError(t, err)
	})

	t.Run("Should return no error adding the first Profile for a User and getting it back", func(t *testing.T) {

		user, _ := domain.NewUser("uniquename123", "unique@email.com", "uniquepassword")
		want, _ := domain.NewProfile("Alias1", 12345, true)
		user.AddProfile(*want)
		got := user.GetProfile(0)

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}		
		
	})

	t.Run("Should return an error setting more than one owner profile", func(t *testing.T) {
		user, _ := domain.NewUser("uniquename123", "unique2345@email.com", "uniquepassword")
		p1, _ := domain.NewProfile("Alias1", 12345, true)
		p2, _ := domain.NewProfile("Alias2", 12345, true)
		user.AddProfile(*p1)
		err := user.AddProfile(*p2)
		assertError(t, err, domain.ErrMorethanOneOwner)
	})


}

