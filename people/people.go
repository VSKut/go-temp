package people

import (
	"fmt"
	"sort"
)

type People []Person

func (pp People) Append(ps Person) People {
	return append(pp, ps)
}

func (pp People) String() (r string) {
	for k, v := range pp {
		r += fmt.Sprintf("%d. %s %s (%s)\n", k+1, v.firstName, v.lastName, v.birthDay)
	}

	return
}

func (pp People) Len() int {
	return len(pp)
}
func (pp People) Swap(i, j int) {
	pp[i], pp[j] = pp[j], pp[i]
}
func (pp People) Less(i, j int) bool {
	if pp[i].birthDay.After(pp[j].birthDay) {
		return true
	}

	s := []string{pp[i].firstName + pp[i].lastName, pp[j].firstName + pp[j].lastName}
	if pp[i].birthDay.Equal(pp[j].birthDay) && sort.StringsAreSorted(s) {
		return true
	}

	return false
}

func (pp People) Sort() {
	sort.Sort(pp)
}
