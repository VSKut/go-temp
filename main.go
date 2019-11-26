package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func multiplyByTwo(k *int) {
	*k *= 2
}

func printMoreTen(k int) error {
	if k < 10 {
		return errors.New("k must be > 10")
	}
	fmt.Println(k)
	return nil
}

type jsStruct struct {
	Data int  `json:"data"`
	OK   bool `json:"ok"`
}

func dejson() (jsStruct, error) {
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	if err := json.Unmarshal(in, &out); err != nil {
		return jsStruct{}, err
	}

	return out, nil
}

func main() {
	var r int = 11
	multiplyByTwo(&r)

	if err := printMoreTen(r); err != nil {
		panic(err)
	}
}
