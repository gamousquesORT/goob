package domain_test

import (
	"goob/app/pkg/domain"
	"testing"
)

var adminUser = domain.UserData{}

func createAdminUser(t testing.TB) error {
	t.Helper()
	adminUser, err := domain.NewUser("Admin6789101", "admin@example.com", "password901")
	adminUser.SetAdmin(true)

	return err

}

func TestValidAdminUser(t *testing.T) {
	t.Run("Should return true given an Admin user ", func(t *testing.T) {
		err := createAdminUser(t)
		adminUser.SetAdmin(true)
		got := adminUser.GetAdmin()
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		assertNoError(t, err)

	})
}
