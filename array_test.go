package gojsmodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	strArr  = []string{"a", "b", "c"}
	intArr  = []int32{1, 1, 1, 1, 1, 2, 3}
	boolArr = []bool{true, true, true, true}
)

func Test_Contains(t *testing.T) {

	res1 := Contains(strArr, "a")
	res2 := Contains(strArr, "z")

	assert.Equal(t, res1, true)
	assert.Equal(t, res2, false)

	res3 := Contains(intArr, 1)
	res4 := Contains(boolArr, false)

	assert.Equal(t, res3, true)
	assert.Equal(t, res4, false)
}
