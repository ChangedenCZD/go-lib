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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestVectorAdd(t *testing.T) {
	vector := NewVector()
	vector.Add(100)
	assert.Equal(t, 1, vector.Size())

	vector.AddElement(200)
	assert.Equal(t, 2, vector.Size())
}

func TestVectorRemove(t *testing.T) {
	vector := NewVector()
	vector.Add(100)
	vector.AddElement(200)
	vector.AddElement(300)

	vector.Remove(1)
	assert.Equal(t, 2, vector.Size())
	item := vector.LastElement()
	assert.Equal(t, 300, item)
	item = vector.FirstElement()
	assert.Equal(t, 100, item)

	vector.RemoveElementAt(0)
	assert.Equal(t, 1, vector.Size())

	vector.RemoveAllElements()
	assert.Equal(t, 0, vector.Size())
}

func TestVectorClone(t *testing.T) {
	vector0 := NewVector()
	for i := 0; i < 100; i++ {
		vector0.Add(i)
	}
	vector1 := vector0.Clone()
	assert.Equal(t, vector0.Size(), vector1.Size())

	enumeration0 := vector0.Elements()
	enumeration1 := vector1.Elements()

	for enumeration0.HasMoreElements() && enumeration1.HasMoreElements() {
		item0 := enumeration0.NextElement()
		item1 := enumeration1.NextElement()
		assert.Equal(t, item0, item1)
	}
}
