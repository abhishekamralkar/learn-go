package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	newRectangle := Rectangle{10, 30}
	fmt.Println(Perimeter(newRectangle))
	fmt.Println(newRectangle.Area())
	newCircle := Circle{20}
	fmt.Println(newCircle.Area())
}
