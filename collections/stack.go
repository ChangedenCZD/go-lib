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

type stack struct {
	*vector
}

func NewStack() *stack {
	inst := &stack{
		vector: NewVector(),
	}
	return inst
}

func (s *stack) Empty() bool {
	return s.Size() == 0
}

func (s *stack) Push(item interface{}) interface{} {
	s.AddElement(item)
	return item
}

func (s *stack) Pop() interface{} {
	obj := s.Peek()
	s.RemoveElementAt(s.Size() - 1)
	return obj
}

func (s *stack) Peek() interface{} {
	l := s.Size()
	if l == 0 {
		panic("empty stack")
	}
	return s.ElementAt(l - 1)
}

func (s *stack) Search(obj interface{}) int {
	i := s.LastIndexOf(obj, s.elementCount-1)

	if i >= 0 {
		return s.Size() - i
	}

	return -1
}
