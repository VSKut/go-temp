package main

import (
	"github.com/vskut/go-temp/printsorted"
)

func main() {

	m := map[int]string{2: "a", 0: "b", 1: "c"}
	printsorted.PrintSorted(m)

	m2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printsorted.PrintSorted(m2)
}
