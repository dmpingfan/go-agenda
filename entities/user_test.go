package entities_test

import (
	"testing"

	"github.com/painterdrown/go-agenda/entities"
)

func init() {
	entities.InitDB("/tmp/go-agenda.db")
}

func TestUser(t *testing.T) {
	var err error
	if err = entities.Register("test", "test"); err != nil {
		t.Error("Register 错误")
	}
	if _, err = entities.GetAllUsers(); err != nil {
		t.Error("GetAllUsers 错误")
	}
	if err = entities.DeleteUser("test"); err != nil {
		t.Error("DeleteUser 错误")
	}
}
