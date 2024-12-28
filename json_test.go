package gojsmodule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Information struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Job string `json:"job"`
}

func getInformation() Information{
	
	return Information{
		Name : "leedonggyu",
		Age: 31,
		Job: "Programmer",
	}
}

func Test_JsonStringify(t *testing.T) {
	p := getInformation()

	b, s, err  := JsonStringify(p)

	assert.Equal(t, err ,nil)
	fmt.Println("buffer : ", b)
	fmt.Println("string : ", s)
}

func Test_JsonParse(t *testing.T) {

	p := getInformation()
	b, _, err := JsonStringify(p)
	assert.Equal(t, err, nil)

	obj := JsonParse[Information](b)
	
	assert.Equal(t, obj.Name, "leedonggyu")
	assert.Equal(t, obj.Age, 31)
	assert.Equal(t, obj.Job, "Programmer")
}