package tools

import (
	"reflect"
)

func UnpackType(v reflect.Type) reflect.Type {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}

func UnpackValue(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}
