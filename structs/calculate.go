package structs

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	w float64
	h float64
}

func (r Rectangle) Area() float64 {
	return r.h * r.w
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.r * c.r)
}

type Triangle struct {
	b float64
	h float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.h)
}
func Perimeter(values Rectangle) float64 {
	return 2 * (values.w + values.h)
}

/*func Area(values Rectangle) float64 {
	return values.w * values.h
}*/
