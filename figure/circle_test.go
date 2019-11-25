package figure

import "testing"

func TestCircle_Area(t *testing.T) {
	c := Circle{Radius: 10}

	if a := c.Area(); a != float64(314.1592653589793) {
		t.Error("Circle{Radius: 10}.Area() != 314.1592653589793. Returned: ", a)
	}
}

func TestCircle_Perimeter(t *testing.T) {
	c := Circle{Radius: 10}

	if a := c.Perimeter(); a != float64(62.83185307179586) {
		t.Error("Circle{Radius: 10}.Perimeter() != 62.83185307179586. Returned: ", a)
	}
}
