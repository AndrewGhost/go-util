package structs

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

//将元素类型为struct的slice转换为map
func convertToMap(s interface{}, key string) map[interface{}]interface{} {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return nil
	}

	if reflect.TypeOf(s).Elem().Kind() != reflect.Struct {
		return nil
	}

	refv := reflect.ValueOf(s)
	if refv.Len() == 0 {
		return nil
	}

	if !refv.Index(0).FieldByName(key).IsValid() {
		return nil
	}

	retMap := make(map[interface{}]interface{})
	for i := 0; i < refv.Len(); i++ {
		field := refv.Index(i).FieldByName(key)
		fmt.Println(field.Interface())
		retMap[field.Interface()] = refv.Index(i).Interface()
	}

	return retMap
}
