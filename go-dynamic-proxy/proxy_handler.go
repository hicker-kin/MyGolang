package proxy

import (
	"log"
	"reflect"
	"time"
)

type ServiceDynamicProxy struct {
	target      interface{}  // proxy struct
	interceptor ProxyHandler // proxy method(not contains function)
}

type ProxyHandler = func(
	method string,
	args []reflect.Value,
	methodFunc func(in []reflect.Value) []reflect.Value, // target method
) []reflect.Value

// Call
/**
 * @author qinzj
 * @description: execute proxy
 * @date 14:55 2023/4/4
 * @param:
 *	methodFunc: the method of receiver, such as: StructA.Method1
 *	args: the reflect value slice of args, such as: []reflect.Value{reflect.ValueOf("proxy message"),}
 * @return: the reflect value of proxy interceptor
 */
func (p *ServiceDynamicProxy) Call(methodFunc interface{}, args []reflect.Value) []reflect.Value {
	methodName := handlerName(methodFunc)

	// The proxy subject executes proxy methods, which include the method
	// being proxied and the user's own enhanced behavior
	return p.interceptor(methodName, args, reflect.ValueOf(p.target).MethodByName(methodName).Call)
}

// BuildServiceDynamicProxyInstance
/**
 * @author qinzj
 * @description: create a proxy subject
 * @date 14:57 2023/4/4
 * @param:
 * target: proxy target
 * h: proxy behavior which include the method of target
 * @return: proxy
 **/
func BuildServiceDynamicProxyInstance(target interface{},
	h ProxyHandler) *ServiceDynamicProxy {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr { // the target type is reflect.Ptr
		panic("need a pointer type of interface struct")
	}

	if h == nil {
		h = defaultAOPHandler
	}
	return &ServiceDynamicProxy{
		target:      target,
		interceptor: h,
	}
	/*proxyValue := reflect.New(reflect.TypeOf(proxy).Elem())

	targetValue := reflect.ValueOf(target)
	proxyValue.Elem().FieldByName("target").Set(targetValue)
	proxyValue.Elem().FieldByName("interceptor").Set(reflect.ValueOf(proxy.interceptor))
	return proxyValue.Interface().(*ServiceDynamicProxy)*/
}

func defaultAOPHandler(
	method string,
	args []reflect.Value,
	methodFunc func(in []reflect.Value) []reflect.Value) []reflect.Value {
	return func(
		method string,
		args []reflect.Value,
		next func(in []reflect.Value) []reflect.Value,
	) []reflect.Value {
		// set your handler
		log.Printf("Beginning call: %s\n", method)

		start := time.Now().UnixMilli()
		result := next(args)
		log.Printf("Finished call: %s, spend {%d} ms\n", method, time.Now().UnixMilli()-start)
		return result
	}(method, args, methodFunc)
}
