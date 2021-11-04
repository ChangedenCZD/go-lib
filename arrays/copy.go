package arrays

import (
	"fmt"
	"math"
	"reflect"
)

func Copy(src interface{}, dst interface{}) {
	CopyOfRange(src, 0, unpackV(reflect.ValueOf(src)).Len(), dst)
}

func CopyOf(src interface{}, newLength int, dst interface{}) {
	if newLength < 0 {
		panic(fmt.Sprintf("negative array size: %d", newLength))
	}
	CopyOfRange(src, 0, newLength, dst)
}

func CopyOfRange(src interface{}, start int, end int, dst interface{}) {
	if start > end {
		panic("illegal argument")
	}
	originalLength := unpackV(reflect.ValueOf(src)).Len()
	if start < 0 || start > originalLength {
		panic("array index out of bounds")
	}
	resultLength := end - start
	copyLength := int(math.Min(float64(resultLength), float64(originalLength-start)))
	CopyFrom(src, start, dst, 0, copyLength)
}

func unpackT(v reflect.Type) reflect.Type {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}

func unpackV(v reflect.Value) reflect.Value {
	for reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	return v
}

func CopyFrom(s interface{}, srcPos int, d interface{}, dstPos int, length int) {
	sTV := unpackT(reflect.TypeOf(s))
	dTV := unpackT(reflect.TypeOf(d))

	if reflect.Slice != sTV.Kind() &&
		reflect.Array != sTV.Kind() {
		panic("src must be type array")
	}

	if reflect.Slice != dTV.Kind() &&
		reflect.Array != dTV.Kind() {
		panic("dest must be type array")
	}

	sET := sTV.Elem()
	dET := dTV.Elem()

	if sET.Kind() != dET.Kind() ||
		sET.Name() != dET.Name() {
		panic(fmt.Sprintf("src and dest element type not match (src: %s, dest: %s)", sET.Kind(), dET.Kind()))
	}

	if srcPos < 0 || dstPos < 0 || length < 0 {
		panic("array index out of bounds")
	}

	sVV := unpackV(reflect.ValueOf(s))
	dVV := unpackV(reflect.ValueOf(d))

	if length+srcPos > sVV.Len() ||
		length+dstPos > dVV.Len() {
		panic("array index out of bounds")
	}

	if length == 0 {
		return
	}

	for sIndex := 0; sIndex < length; sIndex++ {
		dVV.Index(dstPos + sIndex).Set(sVV.Index(srcPos + sIndex))
	}
}
