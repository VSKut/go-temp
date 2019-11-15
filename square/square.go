package square

// Square ...
type Square struct {
	StartPoint Point
	Side       uint
}

// End ...
func (s Square) End() (uint, uint) {
	return s.Side, s.Side
}

// Perimeter ...
func (s Square) Perimeter() uint {
	return s.Side * 4
}

// Area ...
func (s Square) Area() uint {
	return s.Side * s.Side
}
