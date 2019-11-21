package printsorted

import (
	"fmt"
	"sort"
)

func PrintSorted(m map[int]string) {
	for _, v := range sortMap(m) {
		fmt.Print(v, " ")
	}

	fmt.Println()
}

func sortMap(m map[int]string) []string {
	k, r := make([]int, 0, len(m)), make([]string, 0, len(m))
	for index := range m {
		k = append(k, index)
	}
	sort.Ints(k)
	for _, value := range k {
		r = append(r, m[value])
	}
	return r
}
