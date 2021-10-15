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

type stringBuilder struct {
	*abstractStringBuilder
}

func NewStringBuilder() *stringBuilder {
	return NewStringBuilderCap(16)
}

func NewStringBuilderCap(capacity int) *stringBuilder {
	inst := &stringBuilder{
		abstractStringBuilder: NewAbstractStringBuilder(capacity),
	}
	return inst
}

func NewStringBuilderStr(str string) *stringBuilder {
	inst := NewStringBuilderCap(len([]rune(str)) + 16)
	return inst.Append(str)
}

func (b *stringBuilder) Append(v interface{}) *stringBuilder {
	b.abstractStringBuilder.append(v)
	return b
}

func (b *stringBuilder) AppendNil() *stringBuilder {
	b.abstractStringBuilder.appendNil()
	return b
}

func (b *stringBuilder) AppendRune(v rune) *stringBuilder {
	b.abstractStringBuilder.appendRune(v)
	return b
}

func (b *stringBuilder) AppendRuneArray(v []rune) *stringBuilder {
	b.abstractStringBuilder.appendRuneArray(v)
	return b
}

func (b *stringBuilder) AppendString(v string) *stringBuilder {
	b.abstractStringBuilder.appendString(v)
	return b
}

func (b *stringBuilder) AppendObject(v interface{}) *stringBuilder {
	b.abstractStringBuilder.appendObject(v)
	return b
}

func (b *stringBuilder) AppendCodePoint(codePoint int) *stringBuilder {
	b.abstractStringBuilder.appendCodePoint(codePoint)
	return b
}

func (b *stringBuilder) Delete(start int, end int) *stringBuilder {
	b.abstractStringBuilder.delete(start, end)
	return b
}

func (b *stringBuilder) DeleteCharAt(index int) *stringBuilder {
	b.abstractStringBuilder.deleteCharAt(index)
	return b
}

func (b *stringBuilder) Replace(start int, end int, str string) *stringBuilder {
	b.abstractStringBuilder.replace(start, end, str)
	return b
}

func (b *stringBuilder) Insert(index int, v interface{}) *stringBuilder {
	b.abstractStringBuilder.insert(index, v)
	return b
}

func (b *stringBuilder) IndexOf(str string) int {
	return b.IndexOfFrom(str, 0)
}

func (b *stringBuilder) IndexOfFrom(str string, from int) int {
	return b.abstractStringBuilder.indexOf(str, from)
}

func (b *stringBuilder) LastIndexOf(str string) int {
	return b.LastIndexOfFrom(str, b.Length())
}

func (b *stringBuilder) LastIndexOfFrom(str string, from int) int {
	return b.abstractStringBuilder.lastIndexOf(str, from)
}

func (b *stringBuilder) Reverse() *stringBuilder {
	b.abstractStringBuilder.reverse()
	return b
}

func (b *stringBuilder) String() string {
	return string(b.value[0:b.count])
}
