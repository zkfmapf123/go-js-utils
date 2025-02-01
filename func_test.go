package gojsmodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addData(v int) (int, error) {
	return v+v, nil
}

func Test_Pipe(t *testing.T) {
	
	firstValue := 1

	v, err :=Pipe(firstValue,
		addData,
		addData,
		addData,
		addData,
	)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, v, 16)
}