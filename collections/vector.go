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

package collections

import (
	"fmt"
	"sync"
)

type Vector struct {
	sync.Mutex

	elementCount      int
	capacityIncrement int

	elementData []interface{}
}

func NewVector() *Vector {
	inst := &Vector{
		elementData: make([]interface{}, 0),
	}
	return inst
}

func (v *Vector) Size() int {
	return v.elementCount
}

func (v *Vector) Capacity() int {
	return cap(v.elementData)
}

func (v *Vector) IsEmpty() bool {
	return v.elementCount == 0
}

func (v *Vector) AddElement(obj interface{}) {
	v.Lock()
	defer v.Unlock()
	v.elementCount++
	v.elementData = append(v.elementData, obj)
}

func (v *Vector) RemoveElementAt(index int) {
	v.Lock()
	defer v.Unlock()
	if index >= v.elementCount {
		panic(fmt.Sprintf("array index out of bounds (%d >= %d)", index, v.elementCount))
	}
	if index < 0 {
		panic(fmt.Sprintf("array index out of bounds (%d)", index))
	}
	l := v.Size()
	elementData := v.elementData
	if index == 0 {
		if l > 1 {
			v.elementData = elementData[1:]
		} else {
			v.elementData = make([]interface{}, 0)
		}
	} else if index == l-1 {
		v.elementData = elementData[:index]
	} else {
		v.elementData = append(elementData[:index], elementData[index+1:l]...)
	}
	v.elementCount--
}

func (v *Vector) ElementAt(index int) interface{} {
	if index >= v.elementCount {
		panic(fmt.Sprintf("array index out of bounds (%d >= %d)", index, v.elementCount))
	}
	return v.elementData[index]
}

func (v *Vector) FirstElement() interface{} {
	if v.IsEmpty() {
		panic("no such element")
	}
	return v.elementData[0]
}

func (v *Vector) LastElement() interface{} {
	if v.IsEmpty() {
		panic("no such element")
	}
	return v.elementData[v.elementCount-1]
}

func (v *Vector) SetElementAt(obj interface{}, index int) {
	v.Lock()
	defer v.Unlock()
	if index >= v.elementCount {
		panic(fmt.Sprintf("array index out of bounds (%d >= %d)", index, v.elementCount))
	}
	v.elementData[index] = obj
}

func (v *Vector) InsertElementAt(obj interface{}, index int) {
	v.Lock()
	defer v.Unlock()
	if index > v.elementCount {
		panic(fmt.Sprintf("array index out of bounds (%d > %d)", index, v.elementCount))
	}
	l := v.Size()
	elementData := v.elementData
	if index == 0 {
		v.elementData = append([]interface{}{obj}, elementData...)
	} else if index == l-1 {
		v.elementData = append(elementData, obj)
	} else {
		afterPart := elementData[index:l]
		elementData = append(elementData[:index], obj)
		v.elementData = append(elementData, afterPart...)
	}
	v.elementCount++
}

func (v *Vector) RemoveElement(obj interface{}) bool {
	v.Lock()
	i := v.IndexOf(obj, 0)
	v.Unlock()
	if i >= 0 {
		v.RemoveElementAt(i)
		return true
	}
	return false
}

func (v *Vector) RemoveAllElements() {
	v.Lock()
	defer v.Unlock()
	for i := range v.elementData {
		v.elementData[i] = nil
	}
	v.elementData = make([]interface{}, 0)
	v.elementCount = 0
}

func (v *Vector) Clone() *Vector {
	inst := NewVector()
	elementData := make([]interface{}, 0)
	elementData = append(elementData, v.elementData...)
	inst.elementData = elementData
	inst.elementCount = len(elementData)
	return inst
}

func (v *Vector) Contains(obj interface{}) bool {
	return v.IndexOf(obj, 0) >= 0
}

func (v *Vector) IndexOf(obj interface{}, index int) int {
	for i := index; i < v.elementCount; i++ {
		if v.elementData[i] == obj {
			return i
		}
	}
	return -1
}

func (v *Vector) LastIndexOf(obj interface{}, index int) int {
	if index >= v.elementCount {
		panic(fmt.Sprintf("array index out of bounds (%d >= %d)", index, v.elementCount))
	}
	for i := index; i >= 0; i-- {
		if v.elementData[i] == obj {
			return i
		}
	}
	return -1
}

func (v *Vector) Get(index int) interface{} {
	return v.ElementAt(index)
}

func (v *Vector) Set(index int, obj interface{}) interface{} {
	oldValue := v.Get(index)
	v.SetElementAt(obj, index)
	return oldValue
}

func (v *Vector) Add(obj interface{}) bool {
	v.AddElement(obj)
	return true
}

func (v *Vector) Remove(index int) interface{} {
	oldValue := v.Get(index)
	v.RemoveElementAt(index)
	return oldValue
}

func (v *Vector) Elements() enumerationI {
	return NewEnumeration(v)
}
