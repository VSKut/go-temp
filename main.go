package main

import "./average"

func main() {

	var array = []int{1, 2, 3, 4, 5, 6}
	var emptyArray []int

	println(average.Average(array), average.Average(emptyArray))

}
