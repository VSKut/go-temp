package reverse

func Reverse(s []int64) []int64 {
	l := len(s)
	r := make([]int64, l, l)
	for i := l - 1; i >= 0; i-- {
		r[l-i-1] = s[i]
	}

	return r
}
