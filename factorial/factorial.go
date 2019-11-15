package factorial

// Factorial ...
func Factorial(i uint) uint {
	f := uint(1)

	for l := uint(1); l <= i; l++ {
		f *= l
	}

	return f
}
