package factorial

import "testing"

// TestFactorial ...
func TestFactorial(t *testing.T) {

	if r := Factorial(5); r != 120 {
		t.Error("Factorial(5) != \"120\". Returned:", r)
	}

}
