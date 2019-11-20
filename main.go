package main

import "github.com/vskut/go-temp/max"

func main() {

	var s = []string{"one", "two", "three"}
	println(max.Max(s))

	var s2 = []string{"one", "two"}
	println(max.Max(s2))

}
