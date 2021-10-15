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

import "sync"

type enumerationI interface {
	HasMoreElements() bool
	NextElement() interface{}
	Size() int
}

type enumeration struct {
	sync.Mutex

	count  int
	vector *vector
}

func NewEnumeration(vector *vector) *enumeration {
	return &enumeration{
		count:  0,
		vector: vector,
	}
}

func (v *enumeration) HasMoreElements() bool {
	return v.count < v.vector.Size()
}

func (v *enumeration) NextElement() interface{} {
	v.Lock()
	defer v.Unlock()
	if v.HasMoreElements() {
		count := v.count
		v.count++
		return v.vector.ElementAt(count)
	}
	panic("no such element [Vector Enumeration]")
}

func (v *enumeration) Size() int {
	return v.vector.Size()
}
