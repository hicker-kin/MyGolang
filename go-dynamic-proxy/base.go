package proxy

import (
	"reflect"
	"runtime"
	"strings"
)

func handlerName(f interface{}) string {
	// fullName := nameOfFunction(f)
	return methodRealName(nameOfFunction(f))
}

func methodRealName(handleName string) string {
	fnPathName := strings.Split(handleName, ".")
	if len(fnPathName) == 0 {
		return ""
	}
	fnName := fnPathName[len(fnPathName)-1]
	data := strings.Split(fnName, "-")
	if len(data) == 0 {
		return ""
	}
	return data[0]
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
