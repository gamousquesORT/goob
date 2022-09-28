package domain_test

import (
	"goob/domain/cmd/tlmec"
	"testing"
)

var adminUser = domain.UserData{}

func createAdminUser(t testing.TB) {
	t.Helper()
	adminUser, _ := domain.NewUser("Admin6789101", "admin@example.com", "password")
	adminUser.SetAdmin(true)

}

func TestValidAdminUser(t *testing.T) {
	t.Run("Should return true given an Admin user ", func(t *testing.T) {
		createAdminUser(t)
		adminUser.SetAdmin(true)
		got := adminUser.GetAmdin()
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})
}
