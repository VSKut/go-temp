package average

import "testing"

func TestAverage(t *testing.T) {

	var array = []int{1, 2, 3, 4, 5, 6}

	if r := Average(array); r != 3.5 {
		t.Error("Average([]int{1, 2, 3, 4, 5, 6}) != \"3.5\". Returned:", r)
	}

}

func TestAverageEmptyArray(t *testing.T) {

	var array []int

	if r := Average(array); r != 0 {
		t.Error("Average([]int{}) != \"0\". Returned:", r)
	}

}
