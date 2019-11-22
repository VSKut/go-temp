package main

import (
	"fmt"
	"github.com/vskut/go-temp/figure"
)

func main() {
	var s figure.Figure = figure.Square{Width: 10}
	var c figure.Figure = figure.Circle{Radius: 10}

	fmt.Println(s.Area(), s.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())

}
