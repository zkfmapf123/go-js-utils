package gojsmodule

type wrapArr[T any] struct {
	arr []T
}

func WrapArr[T any](arr []T) *wrapArr[T] {

	return &wrapArr[T]{
		arr: arr,
	}
}

func (wr *wrapArr[T]) Map(fn func(value T) T) *wrapArr[T] {
	newArr := []T{}

	for _, v := range wr.arr {

		newArr = append(newArr, fn(v))
	}

	wr.arr = newArr

	return wr
}

func (wr *wrapArr[T]) Filter(fn func(value T) bool) *wrapArr[T] {

	newArr := []T{}

	for _, v := range wr.arr {

		if !fn(v) {
			continue
		}

		newArr = append(newArr, v)
	}

	wr.arr = newArr
	return wr
}

func (wr *wrapArr[T]) Reduce(fn func(sum T, value T) T) *wrapArr[T] {
	var sum T

	for _, v := range wr.arr {
		sum = fn(sum, v)
	}

	wr.arr = []T{sum}
	return wr
}

func (wr *wrapArr[T]) End() []T {
	return wr.arr
}
