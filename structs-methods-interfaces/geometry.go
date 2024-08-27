package geometry

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Height float64
	Width  float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Perimeter(rec Rectangle) float64 {
	return rec.Height*2 + rec.Width*2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
