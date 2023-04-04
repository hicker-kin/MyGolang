package example

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	proxy "github.com/hicker-kin/MyGolang/go-dynamic-proxy"
)

type ServiceInterface interface {
	Method1()
	Method2()
	Method3(msg string, m map[string]string) string
}

type ServiceImpl struct {
	Func func()
}

func (s ServiceImpl) Method1() {
	fmt.Println("In ServiceImpl.Method1")
}

func (s ServiceImpl) Method2() {
	fmt.Println("In ServiceImpl.Method2")
}

func (s ServiceImpl) Method3(msg string, m map[string]string) string {
	fmt.Println("In ServiceImpl.Method3")
	fmt.Printf("message: %s, data: %#+v \n", msg, m)
	return fmt.Sprintf("%s success!", msg)
}

func TestDynamicProxy(t *testing.T) {
	// realTarget := ServiceImpl{ // need a pointer type of interface struct
	realTarget := &ServiceImpl{
		Func: func() {
			fmt.Println("this is inner function: Func")
		},
	}

	// create proxy
	proxy := proxy.BuildServiceDynamicProxyInstance(realTarget, nil)

	// call method
	result := proxy.Call(realTarget.Method1, nil)
	t.Log("result 1----", result)

	result = proxy.Call(realTarget.Method3, []reflect.Value{
		reflect.ValueOf("proxy message"),
		reflect.ValueOf(map[string]string{"key": "val"}),
	})
	t.Log("result 3----", result)
}

func TestDynamicProxyWithCustomHandler(t *testing.T) {
	// real subject
	realTarget := &ServiceImpl{}

	// create proxy
	proxy := proxy.BuildServiceDynamicProxyInstance(realTarget, func(
		method string,
		args []reflect.Value,
		methodFunc func(in []reflect.Value) []reflect.Value) []reflect.Value {
		log.Printf("Custom Before %s\n", method)
		result := methodFunc(args)
		log.Printf("Custom After %s\n", method)
		return result
	})

	// call method
	result := proxy.Call(realTarget.Method1, nil)
	t.Log("result 1----", result)

	result = proxy.Call(realTarget.Method3, []reflect.Value{
		reflect.ValueOf("proxy message"),
		reflect.ValueOf(map[string]string{"key": "val"}),
	})
	t.Log("result 3----", result)
}
