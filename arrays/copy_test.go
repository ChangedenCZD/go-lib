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
