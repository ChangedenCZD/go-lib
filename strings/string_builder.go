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

type StringBuilder struct {
	*abstractStringBuilder
}

func NewStringBuilder() *StringBuilder {
	return NewStringBuilderCap(16)
}

func NewStringBuilderCap(capacity int) *StringBuilder {
	inst := &StringBuilder{
		abstractStringBuilder: NewAbstractStringBuilder(capacity),
	}
	return inst
}

func NewStringBuilderStr(str string) *StringBuilder {
	inst := NewStringBuilderCap(len([]rune(str)) + 16)
	return inst.Append(str)
}

func (b *StringBuilder) Append(v interface{}) *StringBuilder {
	b.abstractStringBuilder.append(v)
	return b
}

func (b *StringBuilder) AppendNil() *StringBuilder {
	b.abstractStringBuilder.appendNil()
	return b
}

func (b *StringBuilder) AppendRune(v rune) *StringBuilder {
	b.abstractStringBuilder.appendRune(v)
	return b
}

func (b *StringBuilder) AppendRuneArray(v []rune) *StringBuilder {
	b.abstractStringBuilder.appendRuneArray(v)
	return b
}

func (b *StringBuilder) AppendString(v string) *StringBuilder {
	b.abstractStringBuilder.appendString(v)
	return b
}

func (b *StringBuilder) AppendObject(v interface{}) *StringBuilder {
	b.abstractStringBuilder.appendObject(v)
	return b
}

func (b *StringBuilder) AppendCodePoint(codePoint int) *StringBuilder {
	b.abstractStringBuilder.appendCodePoint(codePoint)
	return b
}

func (b *StringBuilder) Delete(start int, end int) *StringBuilder {
	b.abstractStringBuilder.delete(start, end)
	return b
}

func (b *StringBuilder) DeleteCharAt(index int) *StringBuilder {
	b.abstractStringBuilder.deleteCharAt(index)
	return b
}

func (b *StringBuilder) Replace(start int, end int, str string) *StringBuilder {
	b.abstractStringBuilder.replace(start, end, str)
	return b
}

func (b *StringBuilder) SetCharAt(index int, str rune) *StringBuilder {
	b.abstractStringBuilder.setCharAt(index, str)
	return b
}

func (b *StringBuilder) Insert(index int, v interface{}) *StringBuilder {
	b.abstractStringBuilder.insert(index, v)
	return b
}

func (b *StringBuilder) IndexOf(str string) int {
	return b.IndexOfFrom(str, 0)
}

func (b *StringBuilder) IndexOfFrom(str string, from int) int {
	return b.abstractStringBuilder.indexOf(str, from)
}

func (b *StringBuilder) LastIndexOf(str string) int {
	return b.LastIndexOfFrom(str, b.Length())
}

func (b *StringBuilder) LastIndexOfFrom(str string, from int) int {
	return b.abstractStringBuilder.lastIndexOf(str, from)
}

func (b *StringBuilder) Reverse() *StringBuilder {
	b.abstractStringBuilder.reverse()
	return b
}

func (b *StringBuilder) String() string {
	return string(b.value[0:b.count])
}
