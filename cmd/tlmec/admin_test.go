package domain_test

import ("testing")


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
