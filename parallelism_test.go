package gojsmodule

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	num = 20000
)

func factorial(n int) (*big.Int, error) {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result, nil
}

func getHeavyJob(isAddErrorFunc bool) []PromiseParams[*big.Int] {
	bp, loop := []PromiseParams[*big.Int]{}, 10

	for i := 0; i < loop; i++ {

		bp = append(bp, PromiseParams[*big.Int]{
			fnName: fmt.Sprintf("factorial-%d", i),
			fn: func() (*big.Int, error) {
				return factorial(num + ((i + 1) * 10000))
			},
		})
	}

	return bp
}

// 순차처리 방식
func executeSimple[T any](fns []PromiseParams[T]) ([]fnResParams[T], int64) {

	start := time.Now()
	fnResult := []fnResParams[T]{}
	for _, attr := range fns {
		res, err, durtaion := execute(attr.fn)
		fnResult = append(fnResult, fnResParams[T]{
			fnName:        attr.fnName,
			result:        res,
			err:           err,
			executionTime: durtaion,
		})
	}

	end := time.Now()
	return fnResult, int64(end.Sub(start).Milliseconds())
}

// ////////////////////////////////////////////// Testing ////////////////////////////////////////////////
func Test_FactorialDuration(t *testing.T) {
	start := time.Now()
	factorial(num)
	end := time.Now()

	duration := end.Sub(start).Milliseconds()
	assert.Equal(t, duration > 1, true) // >= 0.001
}

func Test_benchmarkExecution(t *testing.T) {
	_, simpleExecutionTime := executeSimple(getHeavyJob(false))
	res2, promiseExecutionTime := PromiseAll(getHeavyJob(false))

	assert.Equal(t, simpleExecutionTime > promiseExecutionTime, true)
	for _, attr := range res2 {
		if attr.result.Cmp(big.NewInt(1)) <= 0 {
			t.Errorf("Expected result greater than 1, got: %v", attr.result)
		}
	}
}
