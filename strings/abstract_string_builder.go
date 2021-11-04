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

package strings

import (
	"fmt"
)

import (
	"github.com/changedenczd/go-lib/strings/common"
	error2 "github.com/changedenczd/go-lib/strings/error"
)

const (
	MaxArraySize = common.IntMaxValue - 8
)

type abstractStringBuilder struct {
	value []rune
	count int
}

func NewAbstractStringBuilder(capacity int) *abstractStringBuilder {
	inst := &abstractStringBuilder{
		value: make([]rune, capacity),
	}
	return inst
}

func (s *abstractStringBuilder) Length() int {
	return s.count
}

func (s *abstractStringBuilder) Capacity() int {
	return len(s.value)
}

func (s *abstractStringBuilder) ensureCapacity(minimumCapacity int) {
	if minimumCapacity > 0 {
		s.ensureCapacityInternal(minimumCapacity)
	}
}

func (s *abstractStringBuilder) ensureCapacityInternal(minimumCapacity int) {
	if minimumCapacity-s.Capacity() > 0 {
		s.copyValue(s.newCapacity(minimumCapacity))
	}
}

func (s *abstractStringBuilder) newCapacity(minCapacity int) int {
	newCapacity := (s.Capacity() << 1) + 2
	if newCapacity-minCapacity < 0 {
		newCapacity = minCapacity
	}
	if newCapacity <= 0 || MaxArraySize-newCapacity < 0 {
		return s.hugeCapacity(minCapacity)
	}
	return newCapacity
}

func (s *abstractStringBuilder) hugeCapacity(minCapacity int) int {
	if common.IntMaxValue-minCapacity < 0 { // overflow
		panic("out of memory")
	}
	if minCapacity > MaxArraySize {
		return minCapacity
	}
	return MaxArraySize
}

func (s *abstractStringBuilder) TrimToSize() {
	if s.count < len(s.value) {
		s.copyValue(s.count)
	}
}

func (s *abstractStringBuilder) copyValue(count int) {
	arr := make([]rune, count)
	copy(arr, s.value)
	s.value = arr
}

func (s *abstractStringBuilder) SetLength(newLength int) {
	if newLength < 0 {
		panic(error2.NewStringIndexOutOfBoundsError(newLength))
	}
	s.ensureCapacityInternal(newLength)
	if s.count < newLength {
		for i := s.count; i < newLength; i++ {
			s.value[i] = 0
		}
	}
	s.count = newLength
}

func (s *abstractStringBuilder) CharAt(index int) rune {
	if (index < 0) || (index >= s.count) {
		panic(error2.NewStringIndexOutOfBoundsError(index))
	}
	return s.value[index]
}

func (s *abstractStringBuilder) CodePointAt(index int) int {
	if (index < 0) || (index >= s.count) {
		panic(error2.NewStringIndexOutOfBoundsError(index))
	}
	return codePointAtImpl(s.value, index, s.count)
}

func (s *abstractStringBuilder) CodePointBefore(index int) int {
	i := index - 1
	if (i < 0) || (i >= s.count) {
		panic(error2.NewStringIndexOutOfBoundsError(index))
	}
	return codePointBeforeImpl(s.value, index, 0)
}

func (s *abstractStringBuilder) codePointCount(beginIndex int, endIndex int) int {
	if beginIndex < 0 || endIndex > s.count || beginIndex > endIndex {
		panic(error2.NewIndexOutOfBoundsError(nil))
	}
	return codePointCountImpl(s.value, beginIndex, endIndex-beginIndex)
}

func (s *abstractStringBuilder) appendCodePoint(codePoint int) *abstractStringBuilder {
	count := s.count
	if isBmpCodePoint(codePoint) {
		s.ensureCapacityInternal(count + 1)
		s.value[count] = rune(codePoint)
		s.count = count + 1
	} else if isValidCodePoint(codePoint) {
		s.ensureCapacityInternal(count + 2)
		toSurrogates(codePoint, s.value, count)
		s.count = count + 2
	} else {
		panic(error2.NewIllegalArgumentError(nil))
	}
	return s
}

func (s *abstractStringBuilder) append(v interface{}) *abstractStringBuilder {
	return s.appendRuneArray(getRuneArray(v))
}

func (s *abstractStringBuilder) appendNil() *abstractStringBuilder {
	return s.appendRuneArray(getRuneArray(nil))
}

func (s *abstractStringBuilder) appendRune(v rune) *abstractStringBuilder {
	return s.appendRuneArray(getRuneArray(v))
}

func (s *abstractStringBuilder) appendRuneArray(v []rune) *abstractStringBuilder {
	l := len(v)
	c := s.count
	s.ensureCapacityInternal(c + l)
	for i := 0; i < l; i++ {
		s.value[c+i] = v[i]
	}
	s.count = c + l
	return s
}

func (s *abstractStringBuilder) appendString(v string) *abstractStringBuilder {
	return s.appendRuneArray(getRuneArray(v))
}

func (s *abstractStringBuilder) appendObject(v interface{}) *abstractStringBuilder {
	return s.appendRuneArray(getRuneArray(v))
}

func getRuneArray(v interface{}) []rune {
	switch vv := v.(type) {
	case nil:
		return []rune{'n', 'i', 'l'}
	case rune:
		return []rune{vv}
	case []rune:
		return vv
	case string:
		return []rune(vv)
	default:
		return []rune(fmt.Sprintf("%v", v))
	}
}

func (s *abstractStringBuilder) delete(start int, end int) *abstractStringBuilder {
	count := s.count
	if start < 0 {
		panic(error2.NewStringIndexOutOfBoundsError(start))
	}
	if end > count {
		end = count
	}
	if start > end {
		panic(error2.NewStringIndexOutOfBoundsErrorEmpty())
	}
	l := end - start
	if l > 0 {
		newArr := make([]rune, count-l)
		idx := 0
		for i := 0; i < count; {
			if i < start || i >= end {
				newArr[idx] = s.value[i]
				i++
				idx++
				continue
			}
			i++
			if i == start { // skip to end position
				i = end
			}
		}
		s.value = newArr
		s.count = count - l
	}
	return s
}

func (s *abstractStringBuilder) deleteCharAt(index int) *abstractStringBuilder {
	if (index < 0) || (index >= s.count) {
		panic(error2.NewStringIndexOutOfBoundsError(index))
	}
	count := s.count
	idx := 0
	newArr := make([]rune, count-1)
	for i := 0; i < count; {
		if i == index {
			i++
			continue
		}
		newArr[idx] = s.value[i]
		i++
		idx++
	}
	s.value = newArr
	s.count = count - 1
	return s
}

func (s *abstractStringBuilder) replace(start int, end int, str string) *abstractStringBuilder {
	count := s.count
	if start < 0 {
		panic(error2.NewStringIndexOutOfBoundsError(start))
	}
	if start > count {
		panic(error2.NewStringIndexOutOfBoundsErrorStr("start > length()"))

	}
	if start > end {
		panic(error2.NewStringIndexOutOfBoundsErrorStr("start > end"))
	}
	if end > count {
		end = count
	}
	l := len(str)
	newCount := count + l - (end - start)
	origin := make([]rune, count)
	copy(origin, s.value)
	s.ensureCapacityInternal(newCount)
	strRune := []rune(str)
	originStartPart := origin[0:start]
	originEndPart := origin[end:]
	idx := 0
	for _, v := range originStartPart {
		s.value[idx] = v
		idx++
	}
	for _, v := range strRune {
		s.value[idx] = v
		idx++
	}
	for _, v := range originEndPart {
		s.value[idx] = v
		idx++
	}
	s.count = newCount
	return s
}

func (s *abstractStringBuilder) insert(offset int, v interface{}) *abstractStringBuilder {
	if (offset < 0) || (offset > s.Length()) {
		panic(error2.NewStringIndexOutOfBoundsError(offset))
	}
	count := s.count
	strRune := getRuneArray(v)
	l := len(strRune)
	origin := make([]rune, count)
	copy(origin, s.value)
	s.ensureCapacityInternal(count + l)
	idx := 0
	for i := 0; i < count; i++ {
		if i == offset {
			for _, v := range strRune {
				s.value[idx] = v
				idx++
			}
		}
		s.value[idx] = origin[i]
		idx++

	}
	s.count = count + l
	return s
}

func (s *abstractStringBuilder) indexOf(str string, from int) int {
	target := []rune(str)
	return stringIndexOf(s.value, 0, s.count, target, 0, len(target), from)
}

func (s *abstractStringBuilder) lastIndexOf(str string, from int) int {
	target := []rune(str)
	return stringLastIndexOf(s.value, 0, s.count, target, 0, len(target), from)
}

func (s *abstractStringBuilder) reverse() *abstractStringBuilder {
	hasSurrogates := false
	n := s.count - 1
	for j := (n - 1) >> 1; j >= 0; j-- {
		k := n - j
		cj := s.value[j]
		ck := s.value[k]
		s.value[j] = ck
		s.value[k] = cj
		if isSurrogate(cj) || isSurrogate(ck) {
			hasSurrogates = true
		}
	}
	if hasSurrogates {
		s.reverseAllValidSurrogatePairs()
	}
	return s
}

func (s *abstractStringBuilder) reverseAllValidSurrogatePairs() {
	for i := 0; i < s.count-1; i++ {
		c2 := s.value[i]
		if isLowSurrogate(c2) {
			c1 := s.value[i+1]
			if isHighSurrogate(c1) {
				s.value[i] = c1
				i++
				s.value[i] = c2
			}
		}
	}
}

func (s *abstractStringBuilder) setCharAt(index int, ch rune) *abstractStringBuilder {
	if (index < 0) || (index >= s.count) {
		panic(error2.NewStringIndexOutOfBoundsError(index))
	}
	s.value[index] = ch
	return s
}

func (s *abstractStringBuilder) String() string {
	return string(s.value[0:s.count])
}
