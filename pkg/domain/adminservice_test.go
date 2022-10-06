package domain_test

import (
	"goob/app/pkg/domain"
	"testing"
)


func TestAdminServices(t *testing.T) {
	t.Run("Should create the default Admin User ", func(t *testing.T) {

		admin, err := domain.StreamAppData.CreateAdmin()
		got := admin.IsAdmin
		want := true 
		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}
		assertNoError(t, err)
		
	})


}
