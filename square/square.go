package square

// Square ...
type Square struct {
	StartPoint Point
	Side       uint
}

// End ...
func (s Square) End() Point {
	return Point{
		X: s.StartPoint.X + int(s.Side),
		Y: s.StartPoint.Y + int(s.Side),
	}
}

// Perimeter ...
func (s Square) Perimeter() uint {
	return s.Side * 4
}

// Area ...
func (s Square) Area() uint {
	return s.Side * s.Side
}
