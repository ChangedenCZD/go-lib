package tools

import (
	"reflect"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestUnpackType(t *testing.T) {
	two := 2
	a := &two

	aTyp := reflect.TypeOf(a)
	assert.Equal(t, reflect.Ptr, aTyp.Kind())

	aTyp = UnpackType(aTyp)
	assert.Equal(t, reflect.Int, aTyp.Kind())
}

func TestUnpackValue(t *testing.T) {
	two := 3.14
	a := &two

	aTyp := reflect.ValueOf(a)
	assert.Equal(t, reflect.Ptr, aTyp.Kind())

	aTyp = UnpackValue(aTyp)
	assert.Equal(t, reflect.Float64, aTyp.Kind())
}
