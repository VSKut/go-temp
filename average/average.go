package average

func Average(a []int) float64 {

	if len(a) == 0 {
		return 0
	}

	s := 0
	for _, v := range a {
		s += v
	}

	return float64(s) / float64(len(a))
}
