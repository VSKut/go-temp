package figure

type Square struct {
	Width float64
}

func (s Square) Area() float64 {
	return s.Width * s.Width
}

func (s Square) Perimeter() float64 {
	return s.Width * 4
}
