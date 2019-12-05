package strtoint

import "strconv"

func MyStrToInt2(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int(i), err
}
