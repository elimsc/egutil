package reflectu

import (
	"errors"
	"fmt"
	"reflect"
)

// SetField set object field
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

// Map2Struct map to struct
func Map2Struct(m map[string]interface{}, obj interface{}) error {
	for k, v := range m {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// func main() {
// 	myData := map[string]interface{}{
// 		"Name": "jiang",
// 		"Age":  22,
// 	}

// 	type User struct {
// 		Name string
// 		Age  int
// 	}
// 	var user User
// 	err := Map2Struct(myData, &user)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(user.Name)
// 	fmt.Println(user.Age)
// }
