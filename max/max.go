package max

import "unicode/utf8"

func Max(s []string) (r string) {
	for _, v := range s {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(r) {
			r = v
		}
	}
	return
}
