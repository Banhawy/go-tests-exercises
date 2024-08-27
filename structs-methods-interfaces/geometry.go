package geometry

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func Perimeter(rec Rectangle) float64 {
	return rec.Height*2 + rec.Width*2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
