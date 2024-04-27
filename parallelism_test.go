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

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func getHeavyJob() []PromiseParams[*big.Int] {
	bp := []PromiseParams[*big.Int]{}

	for i := 0; i < 10; i++ {
		bp = append(bp, PromiseParams[*big.Int]{
			fnName: fmt.Sprintf("factorial-%d", i),
			fn: func() *big.Int {
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
		res, durtaion := execute(attr.fn)
		fnResult = append(fnResult, fnResParams[T]{
			fnName:        attr.fnName,
			result:        res,
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

func Test_SimpleRun(t *testing.T) {
	_, simpleJobTotal := executeSimple(getHeavyJob())
	fmt.Println("simple >> ", simpleJobTotal)
}

func Test_PromiseAll(t *testing.T) {
	_, concurrencyJobTotal := PromiseAll(getHeavyJob())
	fmt.Println("concurrency >> ", concurrencyJobTotal)

	// for _, attr := range res {
	// 	if attr.result.Cmp(big.NewInt(1)) <= 0 {
	// 		t.Errorf("Expected result greater than 1, got: %v", attr.result)
	// 	}
	// }
}
