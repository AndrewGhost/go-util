package util

import (
	"runtime"
	"strings"
)

// 获取当前调用函数名字
func GetCurCalleeFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	pathList := strings.Split(name, ".")
	return pathList[len(pathList)-1]
}

// 获取父调用函数名字
func GetParentCallFunc() string {
	pc, _, _, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	pathList := strings.Split(name, ".")
	return pathList[len(pathList)-1]
}
