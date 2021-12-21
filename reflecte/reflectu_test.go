package reflecte_test

import (
	"testing"

	"github.com/elimsc/goe/reflecte"
)

func TestSetField(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	u := User{Name: "u", Age: 100}
	reflecte.SetField(&u, "Age", 20)
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
	reflecte.Map2Struct(u1, &u)
	if u.Name != "u" || u.Age != 20 {
		t.Errorf("TestMap2Struct() failed")
	}
}
