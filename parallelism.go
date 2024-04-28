package gojsmodule

import (
	"sync"
	"time"
)

type PromiseParams[T any] struct {
	fnName string
	fn     func() (T, error)
}

type fnResParams[T any] struct {
	fnName        string
	result        T
	err           error
	executionTime int64 // ms
}

func PromiseAll[T any](fns []PromiseParams[T]) ([]fnResParams[T], int64) {

	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan fnResParams[T])

	fnResult := []fnResParams[T]{}

	for _, attr := range fns {
		wg.Add(1)

		go func(attr PromiseParams[T]) {
			defer wg.Done()
			res, err, durtaion := execute(attr.fn)
			ch <- fnResParams[T]{
				fnName:        attr.fnName,
				result:        res,
				err:           err,
				executionTime: durtaion,
			}
		}(attr)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fnResult = append(fnResult, res)
	}

	end := time.Now()
	return fnResult, int64(end.Sub(start).Milliseconds())
}

func execute[T any](fn func() (T, error)) (T, error, int64) {
	start := time.Now()
	res, err := fn()
	end := time.Now()

	return res, err, int64(end.Sub(start).Milliseconds())
}
