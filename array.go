package gojsmodule

import "reflect"

func Contains[T any](arr []T, target T) bool {

	for _, v := range arr {

		if reflect.DeepEqual(v, target) {
			return true
		}
	}

	return false
}
