package gojsmodule

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	strArr1 = []string{"a", "b", "c", "d"}
	intArr2 = []int{1, 2, 3, 4, 5, 6}
)

func Test_MapStr(t *testing.T) {

	mergeArr := WrapArr(strArr1).Map(func(value string) string {
		return fmt.Sprintf("%s%s", value, value)
	}).Map(func(value string) string {
		return fmt.Sprintf("%s%s", value, "!")
	}).End()

	for _, v := range mergeArr {

		if !strings.Contains(v, "!") {
			log.Fatalf("Not Equels %s", v)
		}
	}
}

func Test_MapReduce(t *testing.T) {
	mergeArr := WrapArr(intArr2).Map(func(value int) int {
		return value * 2
	}).Map(func(value int) int {
		return value * 2
	}).Reduce(func(sum, value int) int {
		return sum + value
	}).End()

	assert.Equal(t, mergeArr[0], 84)
}
