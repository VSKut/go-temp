package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var s = []int64{1, 2, 5, 15}
	var sr = []int64{15, 5, 2, 1}

	if r := Reverse(s); !reflect.DeepEqual(r, sr) {
		t.Error("Reverse([]int64{1, 2, 5, 15}) != \"[]int64{15, 5, 2, 1}\". Returned:", r)
	}
}

func TestReverseInNotNil(t *testing.T) {
	var s = []int64{1, 2, 5, 15}

	if r := Reverse(s); r == nil {
		t.Error("Reverse([]int64{1, 2, 5, 15}) == \"nil\". Returned:", r)
	}
}

func TestReverseIsEmpty(t *testing.T) {
	var s []int64

	if r := Reverse(s); len(r) != 0 {
		t.Error("Reverse([]int64{}) != \"\". Returned:", r)
	}
}
