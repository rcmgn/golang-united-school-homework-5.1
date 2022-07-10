package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (result Point) {
	result.x = s.start.x + int(s.a)
	result.y = s.start.y + int(s.a)
	return
}

func (s Square) Area() (result uint) {
	result = s.a * s.a
	return
}

func (s Square) Perimeter() (result uint) {
	result = s.a * 4
	return
}
