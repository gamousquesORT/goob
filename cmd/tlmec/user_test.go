package domain_test

import (
	"testing"
	"reflect"

	"streamapp.com/domain"
)

func TestUserInit(t *testing.T) {
	t.Run("should return an empty User", func(t *testing.T) {

		got, _ := domain.NewUser("name", "email", "password")

		want := domain.UserData{}
		if reflect.DeepEqual(got,want) {
			t.Errorf("Expected %v user name, got %v", got, want)
		}
	})

}
