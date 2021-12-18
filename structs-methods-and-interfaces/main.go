package main

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	pi := math.Pi
	return math.Round(pi * c.Radius * c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Width) * 2
}

func main() {}
