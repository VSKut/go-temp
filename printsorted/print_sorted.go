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

func sortMap(m map[int]string) map[int]string {
	r := map[int]string{}
	k := make([]int, 0, len(m))
	for index := range m {
		k = append(k, index)
	}
	sort.Ints(k)
	for _, index := range k {

		r[index] = m[index]
	}
	return r
}
