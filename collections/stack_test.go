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

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func TestStackPop(t *testing.T) {
	stack := NewStack()
	stack.Push(2)
	pop := stack.Pop()
	assert.Equal(t, 2, pop)
	assert.Equal(t, 0, stack.Size())
}

func TestStackPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(4)
	peek := stack.Peek()
	assert.Equal(t, 4, peek)
	assert.Equal(t, 1, stack.Size())
}

func TestStackSearch(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	index := stack.Search(5)
	assert.Equal(t, 5, index)
	assert.Equal(t, 10, stack.Size())
}
