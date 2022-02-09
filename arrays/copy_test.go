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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestCopyOf(t *testing.T) {
	src := []string{"1", "2", "4", "8"}
	dst0 := make([]string, len(src))
	CopyOf(src, len(src), dst0)
	assert.Equal(t, src[2], dst0[2])

	dst1 := make([]string, len(src))
	CopyOf(&src, len(src), &dst1)
	assert.Equal(t, src[3], dst1[3])
}

func TestCopyOfRange(t *testing.T) {
	src := []interface{}{"1", 2, 3.14, false, func() {}}
	dest := make([]interface{}, 3)
	CopyOfRange(src, 2, 5, dest)
	assert.Equal(t, src[3], dest[1])
}

func TestCopyAll(t *testing.T) {
	src := []interface{}{"1", 2, 3.14, false}
	dest := make([]interface{}, len(src))
	Copy(src, dest)
	for i := range src {
		assert.Equal(t, src[i], dest[i])
	}
}

func TestCopyFrom(t *testing.T) {
	src := []interface{}{"1", 2, 3.14, false, func() {}}
	dest := make([]interface{}, len(src))
	CopyFrom(src, 1, dest, 2, 3)
	j := 1
	for i := 2; i < len(dest); i++ {
		assert.Equal(t, src[j], dest[i])
		j++
	}
}
