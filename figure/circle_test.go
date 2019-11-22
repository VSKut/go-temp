package figure

import "testing"

func TestCircle_Area(t *testing.T) {
	c := Circle{Radius: 10}

	if a := c.Area(); a != float64(314.1592741012573) {
		t.Error("Circle{Radius: 10}.Area() != 314.1592741012573. Returned: ", a)
	}
}

func TestCircle_Perimeter(t *testing.T) {
	c := Circle{Radius: 10}

	if a := c.Perimeter(); a != float64(62.831854820251465) {
		t.Error("Circle{Radius: 10}.Perimeter() != 62.831854820251465. Returned: ", a)
	}
}
