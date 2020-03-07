package reflectu_test

import (
	"testing"

	"github.com/elimsc/egu/reflectu"
)

func TestSetField(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	u := User{Name: "u", Age: 100}
	reflectu.SetField(&u, "Age", 20)
	if u.Age != 20 {
		t.Errorf("TestSetField() failed")
	}
}

func TestMap2Struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	u1 := map[string]interface{}{
		"Name": "u",
		"Age":  20,
	}
	var u User
	reflectu.Map2Struct(u1, &u)
	if u.Name != "u" || u.Age != 20 {
		t.Errorf("TestMap2Struct() failed")
	}
}
