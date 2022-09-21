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

		got, err := domain.NewUser("uniquename123", "uniqueemail", "uniquepassword")

		want := domain.UserData{"uniquename123", "uniqueemail", "uniquepassword", false}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	})

	t.Run("Should return an error given an exitent email", func(t *testing.T) {
		domain.NewUser("uniquename123", "gaston1@example.com", "uniquepassword")
		_, err := domain.NewUser("uniquename222", "gaston1@example.com", "12345678901")

		assertError(t, err, domain.ErrExistingEmail)

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
