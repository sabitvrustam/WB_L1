package units

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {

	return &Point{
		x: x,
		y: y,
	}
}

func distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func Unit24() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	println(distance(*p1, *p2))
}
