package figure

import "testing"

func TestSquare_Area(t *testing.T) {
	var c Figure = Square{Width: 10}

	if a := c.Area(); a != float64(100) {
		t.Error("Square{Width: 10}.Area() != 100. Returned: ", a)
	}
}

func TestSquare_Perimeter(t *testing.T) {
	var c Figure = Square{Width: 10}

	if a := c.Perimeter(); a != float64(40) {
		t.Error("Square{Width: 10}.Perimeter() != 40. Returned: ", a)
	}
}
