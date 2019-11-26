package strtoint

import (
	"fmt"
)

func MyStrToInt3(s string) (i int, err error) {
	_, err = fmt.Sscanf(s, "%d", &i)
	return
}
