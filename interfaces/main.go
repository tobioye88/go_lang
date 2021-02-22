package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	x, y, radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}


func main() {
	circle := Circle{ radius: 7, x: 0, y: 0 }
	rectangle := Rectangle{ width: 10, height: 5 }

	fmt.Printf("Area of circle is %f \n", getArea(circle))
	fmt.Printf("Area of rectangle is %f \n", getArea(rectangle))
}