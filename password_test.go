package goe_test

import (
	"testing"

	"github.com/elimsc/goe"
)

func TestPasswordUtil(t *testing.T) {
	password := "123456"
	hashedPassword, _ := goe.HashPassword(password)
	if !goe.CheckPasswordHash(password, hashedPassword) {
		t.Errorf("TestPasswordUtil() failed")
	}
}
