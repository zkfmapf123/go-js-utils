package gojsmodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	boolObj = map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}

	strObj = map[string]string{
		"a": "true",
		"b": "true",
		"c": "true",
	}

	intObj = map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}
)

func Test_OKeysObject(t *testing.T) {

	res := OKeys(strObj)
	assert.Equal(t, len(res), 3)
}

func Test_OKeysNotExistObj(t *testing.T) {
	obj := map[string]int{}

	res := OKeys(obj)

	assert.Equal(t, len(res), 0)

}

func Test_OValues(t *testing.T) {

	boolObjRes := OValues(boolObj)
	strObjRes := OValues(strObj)
	intObjRes := OValues(intObj)

	assert.Equal(t, len(boolObjRes), 3)
	assert.Equal(t, len(strObjRes), 3)
	assert.Equal(t, len(intObjRes), 3)

}

func Test_OEntires(t *testing.T) {
	res1 := OEntries(boolObj)
	res2 := OEntries(strObj)
	res3 := OEntries(intObj)

	// expected1 := [][]interface{}{{"a", true}, {"b", true}, {"c", true}}
	// expected2 := [][]interface{}{{"a", "true"}, {"b", "true"}, {"c", "true"}}
	// expected3 := [][]interface{}{{"a", 10}, {"b", 20}, {"c", 30}}

	assert.Equal(t, len(res1), 3)
	assert.Equal(t, len(res2), 3)
	assert.Equal(t, len(res3), 3)
}
