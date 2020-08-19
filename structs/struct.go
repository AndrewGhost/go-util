package structs

import (
	"errors"
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

// 将元素类型为struct的slice转换为map
func ConvertToMap(s interface{}, key string) (map[interface{}]interface{}, error) {
	var (
		isElemPtr bool
	)
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return nil, errors.New("interface not slice type")
	}

	elem := reflect.TypeOf(s).Elem()
	if elem.Kind() == reflect.Ptr {
		isElemPtr = true
		elem = elem.Elem()
	}

	if elem.Kind() != reflect.Struct {
		return nil, errors.New("elem not struct type")
	}

	if _, ok := elem.FieldByName(key); !ok {
		return nil, errors.New("field not exist, field=" + key)
	}

	refv := reflect.ValueOf(s)
	if refv.Len() == 0 {
		return nil, nil
	}

	retMap := make(map[interface{}]interface{})
	for i := 0; i < refv.Len(); i++ {
		var field reflect.Value
		if isElemPtr {
			field = refv.Index(i).Elem().FieldByName(key)

		} else {
			field = refv.Index(i).FieldByName(key)
		}

		retMap[field.Interface()] = refv.Index(i).Interface()
	}

	return retMap, nil
}
