package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShape(s Shape) {
	fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rectangle := Rectangle{Width: 6.4, Height: 6.6}
	circle := Circle{Radius: 5.78}

	PrintShape(rectangle)
	PrintShape(circle)
}
