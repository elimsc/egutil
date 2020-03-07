package egu_test

import (
	"testing"

	"github.com/elimsc/egu"
)

func TestPasswordUtil(t *testing.T) {
	password := "123456"
	hashedPassword, _ := egu.HashPassword(password)
	if !egu.CheckPasswordHash(password, hashedPassword) {
		t.Errorf("TestPasswordUtil() failed")
	}
}
