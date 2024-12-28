package gojsmodule

import "encoding/json"

// JSON.stringify
func JsonStringify[T any] (v T) ([]byte, string, error){

	b, err := json.Marshal(v)
	if err != nil{
		return nil,"",err
	}

	return b, string(b), nil
}

// JSON.parse
func JsonParse[T any](v []byte) T {

	var format T
	json.Unmarshal(v, &format)
	return format
}