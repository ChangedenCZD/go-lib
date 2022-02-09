/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package arrays

import (
	"fmt"
	"math"
	"reflect"
)

import (
	"github.com/changedenczd/go-lib/tools"
)

func Copy(src interface{}, dst interface{}) {
	CopyOfRange(src, 0, tools.UnpackValue(reflect.ValueOf(src)).Len(), dst)
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
	originalLength := tools.UnpackValue(reflect.ValueOf(src)).Len()
	if start < 0 || start > originalLength {
		panic("array index out of bounds")
	}
	resultLength := end - start
	copyLength := int(math.Min(float64(resultLength), float64(originalLength-start)))
	CopyFrom(src, start, dst, 0, copyLength)
}

func CopyFrom(s interface{}, srcPos int, d interface{}, dstPos int, length int) {
	sTV := tools.UnpackType(reflect.TypeOf(s))
	dTV := tools.UnpackType(reflect.TypeOf(d))

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

	sVV := tools.UnpackValue(reflect.ValueOf(s))
	dVV := tools.UnpackValue(reflect.ValueOf(d))

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
