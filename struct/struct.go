package _struct

import (
	"fmt"
	"reflect"
)

// s 是struct指针
// fieldValues的key是struct成员名， value是成员目标修改值
func SetStructFields(s interface{}, fieldValues map[string]interface{}) error {
	value := reflect.ValueOf(s)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("modify object is not pointer type")
	}

	elem := value.Elem()

	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("elem is not struct type")
	}

	for fName, fValue := range fieldValues {
		field := elem.FieldByName(fName)
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.Int64:
				field.SetInt(fValue.(int64))
			case reflect.Uint64:
				field.SetUint(fValue.(uint64))
			case reflect.String:
				field.SetString(fValue.(string))
			case reflect.Float64:
				field.SetFloat(fValue.(float64))
			case reflect.Bool:
				field.SetBool(fValue.(bool))
			default:
			}
		}
	}

	return nil
}
