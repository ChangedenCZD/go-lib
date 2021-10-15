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
	"github.com/changedenczd/go-lib/strings/common"
	merror "github.com/changedenczd/go-lib/strings/error"
)

func codePointAt(a []rune, index int, limit int) int {
	if index >= limit || limit < 0 || limit > len(a) {
		panic(merror.NewIndexOutOfBoundsError(nil))
	}
	return codePointAtImpl(a, index, limit)
}

func codePointAtImpl(a []rune, index int, limit int) int {
	c1 := a[index]
	index++
	if isHighSurrogate(c1) && index < limit {
		c2 := a[index]
		if isLowSurrogate(c2) {
			return toCodePoint(c1, c2)
		}
	}
	return int(c1)
}

func isHighSurrogate(ch rune) bool {
	return ch >= common.MinHighSurrogate && ch < (common.MaxHighSurrogate+1)
}

func isLowSurrogate(ch rune) bool {
	return ch >= common.MinLowSurrogate && ch < (common.MaxLowSurrogate+1)
}

func toCodePoint(high rune, low rune) int {
	return int(((high << 10) + low) + (common.MinSupplementaryCodePoint - (common.MinHighSurrogate << 10) - common.MinLowSurrogate))
}

func codePointBefore(a []rune, index int, start int) int {
	if index <= start || start < 0 || start >= len(a) {
		panic(merror.NewIndexOutOfBoundsError(nil))
	}
	return codePointBeforeImpl(a, index, start)
}

func codePointBeforeImpl(a []rune, index int, start int) int {
	index--
	c2 := a[index]
	if isLowSurrogate(c2) && index > start {
		index--
		c1 := a[index]
		if isHighSurrogate(c1) {
			return toCodePoint(c1, c2)
		}
	}
	return int(c2)
}

func codePointCount(a []rune, offset int, count int) int {
	if count > len(a)-offset || offset < 0 || count < 0 {
		panic(merror.NewIndexOutOfBoundsError(nil))
	}
	return codePointCountImpl(a, offset, count)
}

func codePointCountImpl(a []rune, offset int, count int) int {
	endIndex := offset + count
	n := count
	for i := offset; i < endIndex; {
		i0 := i
		i++
		if isHighSurrogate(a[i0]) && i < endIndex && isLowSurrogate(a[i]) {
			n--
			i++
		}
	}
	return n
}

func isBmpCodePoint(codePoint int) bool {
	return uint(codePoint)>>16 == 0
}

func isValidCodePoint(codePoint int) bool {
	plane := uint(codePoint) >> 16
	mcp := uint(common.MaxCodePoint)
	return plane < ((mcp + 1) >> 16)
}

func toSurrogates(codePoint int, dst []rune, index int) {
	dst[index+1] = lowSurrogate(codePoint)
	dst[index] = highSurrogate(codePoint)
}

func highSurrogate(codePoint int) rune {
	c := uint(codePoint) >> 10
	mscp := uint(common.MinSupplementaryCodePoint) >> 10
	return rune(c + (uint(common.MinHighSurrogate) - mscp))
}

func lowSurrogate(codePoint int) rune {
	return rune(codePoint&0x3ff) + common.MinLowSurrogate
}

func stringIndexOf(source []rune, sourceOffset int, sourceCount int, target []rune, targetOffset int, targetCount int, fromIndex int) int {
	if fromIndex >= sourceCount {
		if targetCount == 0 {
			return sourceCount
		}
		return -1
	}
	if fromIndex < 0 {
		fromIndex = 0
	}
	if targetCount == 0 {
		return fromIndex
	}

	first := target[targetOffset]
	max := sourceOffset + (sourceCount - targetCount)

	for i := sourceOffset + fromIndex; i <= max; i++ {
		/* Look for first character. */
		if source[i] != first {
			i++
			for i <= max && source[i] != first {
				i++
			}
		}

		/* Found first character, now look at the rest of v2 */
		if i <= max {
			j := i + 1
			end := j + targetCount - 1
			for k := targetOffset + 1; j < end && source[j] == target[k]; {
				j++
				k++
			}

			if j == end {
				/* Found whole string. */
				return i - sourceOffset
			}
		}
	}
	return -1
}

func stringLastIndexOf(source []rune, sourceOffset int, sourceCount int, target []rune, targetOffset int, targetCount int, fromIndex int) int {
	rightIndex := sourceCount - targetCount
	if fromIndex < 0 {
		return -1
	}
	if fromIndex > rightIndex {
		fromIndex = rightIndex
	}
	/* Empty string always matches. */
	if targetCount == 0 {
		return fromIndex
	}

	strLastIndex := targetOffset + targetCount - 1
	strLastChar := target[strLastIndex]
	min := sourceOffset + targetCount - 1
	i := min + fromIndex

startSearchForLastChar:
	for {
		for i >= min && source[i] != strLastChar {
			i--
		}
		if i < min {
			return -1
		}
		j := i - 1
		start := j - (targetCount - 1)
		k := strLastIndex - 1
		for j > start {
			if source[j] != target[k] {
				k--
				j--
				i--
				continue startSearchForLastChar
			}
		}
		return start - sourceOffset + 1
	}
}

func isSurrogate(ch rune) bool {
	return ch >= common.MinHighSurrogate && ch < (common.MaxLowSurrogate+1)
}
