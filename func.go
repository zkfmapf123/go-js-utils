package gojsmodule

func Pipe[T any](v T, fns ...func(v T) (T, error)) (T, error) {
	
	value := v

	for _, fn := range fns {
		av, err := fn(value)
		if err != nil {
			return v, err
		}

		value = av
	}

	return value, nil
}

func MustPipe[T any](v T, fns ...func(v T) (T, error)) (T, error) {
	value := v

	result, err := Pipe(value, fns...)
	if err != nil {
		panic(err)
	}

	return result, nil
}
