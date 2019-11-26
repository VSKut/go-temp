package main

import (
	"fmt"
	"github.com/vskut/go-temp/strtoint"
)

func main() {
	res, err := strtoint.MyStrToInt("12345")
	fmt.Println(res, err)

	res2, err2 := strtoint.MyStrToInt2("1125235")
	fmt.Println(res2, err2)

	res3, err3 := strtoint.MyStrToInt3("123456")
	fmt.Println(res3, err3)
}
