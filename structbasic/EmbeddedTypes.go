package structbasic

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type NamedPoint struct {
	Point
	Name string
}
