package main

import (
	"./factorial"
	"./square"
	"fmt"
)

func main() {
	fmt.Println(factorial.Factorial(5))

	s := square.Square{
		StartPoint: square.Point{X: 1, Y: 1},
		Side:       150,
	}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
