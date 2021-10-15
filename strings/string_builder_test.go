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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestStringBuilder(t *testing.T) {
	stringBuilder := NewStringBuilder()
	stringBuilder.Append("test")
	assert.Equal(t, "test", stringBuilder.String())

	stringBuilder.Append(123)
	assert.Equal(t, "test123", stringBuilder.String())

	stringBuilder.Append('中')
	assert.Equal(t, "test123中", stringBuilder.String())
}

func TestStringBuilderByStr(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	assert.Equal(t, "stringBuilder", stringBuilder.String())

	stringBuilder.Append("test")
	assert.Equal(t, "stringBuildertest", stringBuilder.String())

	stringBuilder.Append(123)
	assert.Equal(t, "stringBuildertest123", stringBuilder.String())

	stringBuilder.Append('中')
	assert.Equal(t, "stringBuildertest123中", stringBuilder.String())
}

func TestStringBuilderIndexOf(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	assert.Equal(t, 3, stringBuilder.IndexOf("i"))
	assert.Equal(t, 8, stringBuilder.IndexOfFrom("i", 4))
	assert.Equal(t, 8, stringBuilder.LastIndexOf("i"))
	assert.Equal(t, 3, stringBuilder.LastIndexOfFrom("i", 7))
}

func TestStringBuilderRemove(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	stringBuilder.Delete(1, 3)
	assert.Equal(t, "singBuilder", stringBuilder.String())

	stringBuilder.DeleteCharAt(4)
	assert.Equal(t, "singuilder", stringBuilder.String())
}

func TestStringBuilderReplace(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	stringBuilder.Replace(6, 11, "Join")
	assert.Equal(t, "stringJoiner", stringBuilder.String())
}

func TestStringBuilderInsert(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	stringBuilder.Insert(6, "My")
	assert.Equal(t, "stringMyBuilder", stringBuilder.String())
}

func TestStringBuilderReverse(t *testing.T) {
	stringBuilder := NewStringBuilderStr("stringBuilder")
	stringBuilder.Reverse()
	assert.Equal(t, "redliuBgnirts", stringBuilder.String())
}
