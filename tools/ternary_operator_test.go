package tools

import (
	"math/rand"
	"testing"
	"time"
)

import (
	"github.com/stretchr/testify/assert"
)

func randomCondition() bool {
	t := time.Now().UnixNano() / 1e6
	return t%2 == 1
}

func TestTernaryOperator(t *testing.T) {
	condition := randomCondition()
	trueResult := "true"
	falseResult := "false"
	var actual string
	if condition {
		actual = trueResult
	} else {
		actual = falseResult
	}
	result := TernaryOperator(condition, trueResult, falseResult)
	assert.Equal(t, actual, result)
}

func getNestRandomCondition(f int64, n int32) bool {
	rand.Seed(time.Now().UnixNano() / f)
	randomNum := time.Duration(rand.Int31n(n))
	time.Sleep(randomNum * time.Millisecond)
	return randomCondition()
}

func TestTernaryOperatorNest(t *testing.T) {
	var condition0 bool
	var condition1 bool
	var condition2 bool

	trueResult0 := "true"
	falseResult0 := "false"

	trueResult1 := 48
	falseResult1 := 3.14

	result := TernaryOperatorNest(
		func() bool {
			condition0 = getNestRandomCondition(1e6, 2)
			return condition0
		},
		trueResult0,
		TernaryOperatorNest(
			func() bool {
				condition1 = getNestRandomCondition(1e7, 4)
				return condition1
			},
			trueResult1,
			TernaryOperatorNest(
				func() bool {
					condition2 = getNestRandomCondition(1e6, 6)
					return condition2
				},
				falseResult0,
				falseResult1,
			),
		),
	).Result()

	if condition0 {
		assert.Equal(t, trueResult0, result)
		return
	}
	if condition1 {
		assert.Equal(t, trueResult1, result)
		return
	}
	if condition2 {
		assert.Equal(t, falseResult0, result)
		return
	}
	assert.Equal(t, falseResult1, result)
}
