package main

import (
	"reflect"
)

type NativeMethodMap map[string]func(...Object) Object

type NativeObject struct {
	nativeValue interface{}
	methods     NativeMethodMap
}

var _ Object = &NativeObject{}

func (n *NativeObject) Send(name string, args ...Object) Object {
	return n.methods[name](args...)
}

func NewNativeObject(val interface{}) *NativeObject {
	methods := make(NativeMethodMap)

	self := reflect.ValueOf(val)

	rType := self.Type()
	numMethods := rType.NumMethod()
	for i := 0; i < numMethods; i++ {
		method := rType.Method(i)
		methods[method.Name] = func(args ...Object) Object {
			rArgs := make([]reflect.Value, len(args)+1)
			rArgs[0] = self
			for i, arg := range args {
				rArgs[i+1] = reflect.ValueOf(arg)
			}
			return method.Func.Call(rArgs)[0].Interface().(Object)
		}
	}

	return &NativeObject{
		nativeValue: val,
		methods:     methods,
	}
}
