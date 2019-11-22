package figure

import (
	"github.com/google/gxui/math"
)

type Circle struct {
	Radius float64
}

func (s Circle) Area() float64 {
	return float64(math.Pi) * s.Radius * s.Radius
}

func (s Circle) Perimeter() float64 {
	return float64(math.TwoPi) * s.Radius
}
