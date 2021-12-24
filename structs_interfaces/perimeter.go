package structs_interfaces

import "math"

// interface resolution is implicit, so Rectangle & Circle are Shape
type Shape interface {
	Perimeter() (float64, error)
	Area() (float64, error)
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() (float64, error) {
	return (r.width + r.height) * 2, nil
}

func (r Rectangle) Area() (float64, error) {
	return r.width * r.height, nil
}

func (c Circle) Area() (float64, error) {
	return c.radius * c.radius * math.Pi, nil
}

func (c Circle) Perimeter() (float64, error) {
	return 2 * c.radius * math.Pi, nil
}
