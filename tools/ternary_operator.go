package tools

func TernaryOperator(condition bool, trueResult interface{}, falseResult interface{}) interface{} {
	if condition {
		return trueResult
	}
	return falseResult
}

func TernaryOperatorInt(condition bool, trueResult int, falseResult int) int {
	return TernaryOperator(condition, trueResult, falseResult).(int)
}

func TernaryOperatorInt8(condition bool, trueResult int8, falseResult int8) int8 {
	return TernaryOperator(condition, trueResult, falseResult).(int8)
}

func TernaryOperatorInt16(condition bool, trueResult int16, falseResult int16) int16 {
	return TernaryOperator(condition, trueResult, falseResult).(int16)
}

func TernaryOperatorInt32(condition bool, trueResult int32, falseResult int32) int32 {
	return TernaryOperator(condition, trueResult, falseResult).(int32)
}

func TernaryOperatorInt64(condition bool, trueResult int64, falseResult int64) int64 {
	return TernaryOperator(condition, trueResult, falseResult).(int64)
}

func TernaryOperatorUInt(condition bool, trueResult uint, falseResult uint) uint {
	return TernaryOperator(condition, trueResult, falseResult).(uint)
}

func TernaryOperatorUInt8(condition bool, trueResult uint8, falseResult uint8) uint8 {
	return TernaryOperator(condition, trueResult, falseResult).(uint8)
}

func TernaryOperatorUInt16(condition bool, trueResult uint16, falseResult uint16) uint16 {
	return TernaryOperator(condition, trueResult, falseResult).(uint16)
}

func TernaryOperatorUInt32(condition bool, trueResult uint32, falseResult uint32) uint32 {
	return TernaryOperator(condition, trueResult, falseResult).(uint32)
}

func TernaryOperatorUInt64(condition bool, trueResult uint64, falseResult uint64) uint64 {
	return TernaryOperator(condition, trueResult, falseResult).(uint64)
}

func TernaryOperatorFloat32(condition bool, trueResult float32, falseResult float32) float32 {
	return TernaryOperator(condition, trueResult, falseResult).(float32)
}

func TernaryOperatorFloat64(condition bool, trueResult float64, falseResult float64) float64 {
	return TernaryOperator(condition, trueResult, falseResult).(float64)
}

func TernaryOperatorString(condition bool, trueResult string, falseResult string) string {
	return TernaryOperator(condition, trueResult, falseResult).(string)
}

type TernaryOperatorNestWrap struct {
	condition   TernaryOperatorNestConditionFunc
	trueResult  interface{}
	falseResult interface{}
}

type TernaryOperatorNestConditionFunc func() bool

func TernaryOperatorNest(condition TernaryOperatorNestConditionFunc, trueResult interface{}, falseResult interface{}) *TernaryOperatorNestWrap {
	return &TernaryOperatorNestWrap{
		condition:   condition,
		trueResult:  trueResult,
		falseResult: falseResult,
	}
}

func resolveTernaryOperatorNestResult(result interface{}) interface{} {
	if r, ok := result.(*TernaryOperatorNestWrap); ok {
		return r.Result()
	}
	return result
}

func (t *TernaryOperatorNestWrap) Result() interface{} {
	if t.condition == nil {
		return nil
	}
	return TernaryOperator(
		t.condition(),
		func() interface{} {
			return resolveTernaryOperatorNestResult(t.trueResult)
		},
		func() interface{} {
			return resolveTernaryOperatorNestResult(t.falseResult)
		}).(func() interface{})()
}
