package max

import "testing"

func TestMax(t *testing.T) {

	var s = []string{"one", "two", "three"}

	if r := Max(s); r != "three" {
		t.Error("Average([]string{\"one\", \"two\", \"three\"}) != \"three\". Returned:", r)
	}

}

func TestMaxReturnFirstMatch(t *testing.T) {

	var s = []string{"one", "two", "match", "three"}

	if r := Max(s); r != "match" {
		t.Error("Average([]string{\"one\", \"two\", \"match\", \"three\"}) != \"three\". Returned:", r)
	}

}

func TestMaxEmpty(t *testing.T) {

	var s []string

	if r := Max(s); r != "" {
		t.Error("Average([]string{}) != \"\". Returned:", r)
	}

}
