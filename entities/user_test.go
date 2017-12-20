package entities

import (
	"testing"
)

func TestUser(t *testing.T) {
	var err error
	if err = Register("test", "test"); err != nil {
		t.Error("Register 错误")
	}
	if _, err = GetAllUsers(); err != nil {
		t.Error("GetAllUsers 错误")
	}
	if err = DeleteUser("test"); err != nil {
		t.Error("DeleteUser 错误")
	}
}
