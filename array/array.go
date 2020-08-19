package array

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// target：待验证元素，arr:元素slice
func InArray(target interface{}, arr interface{}) bool {
	s := reflect.ValueOf(arr)

	for i := 0; i < s.Len(); i++ {
		if target == s.Index(i).Interface() {
			return true
		}
	}

	return false
}

// 两个数组的交集
func Intersect(a, b, refRet interface{}) error {
	var (
		ifSlice []interface{}
	)

	m := make(map[interface{}]bool)
	s1 := reflect.ValueOf(a)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()] = true
	}

	s2 := reflect.ValueOf(b)

	for i := 0; i < s2.Len(); i++ {
		if _, ok := m[s2.Index(i).Interface()]; ok {
			ifSlice = append(ifSlice, s2.Index(i).Interface())
		}
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

//两个数组的差集
// eg1. X : []int64{1,2,3,4}, Y: []int64{2,3,5,6}, 差集：[]int64{1,4}
// eg1. X : []int64{2,3,5,6}, Y; []int64{1,2,3,4}, 差集：
func Diff(X, Y, refRet interface{}) error {
	var (
		ifSlice []interface{}
		m       = make(map[interface{}]int)
	)

	s1 := reflect.ValueOf(Y)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()]++
	}

	s2 := reflect.ValueOf(X)
	for i := 0; i < s2.Len(); i++ {
		y := s2.Index(i).Interface()
		if m[y] > 0 {
			m[y]--
			continue
		}
		ifSlice = append(ifSlice, y)
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

// 数组去重
func Unique(slice interface{}, refRet interface{}) error {
	var (
		ifSlice []interface{}
	)

	keys := make(map[interface{}]bool)
	refVal := reflect.ValueOf(slice)

	for i := 0; i < refVal.Len(); i++ {
		entry := refVal.Index(i).Interface()
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			ifSlice = append(ifSlice, entry)
		}
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

// 数组拼装为字符串
func Explode(delimiter string, array interface{}) string {
	var (
		joinStr string
	)

	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return ""
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return ""
	}

	for i := 0; i < s.Len(); i++ {
		joinStr += fmt.Sprintf("%v", s.Index(i).Interface()) + delimiter
	}

	return joinStr[0 : len(joinStr)-1]
}
