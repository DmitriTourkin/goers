package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius
}

func calculateShape(s Shape) float64 {
	return s.Area()
}

func describeValue(t interface{}) { // VERY BROAD IMPLEMENTATION
	fmt.Printf("Type %T, Value: %v", t, t)
}

func main() {
	circle := Circle{5.9}
	rectangle := Rectangle{5.7, 6.7}

	calculateShape(circle)
	calculateShape(rectangle)

	describeValue(circle)
	describeValue(rectangle)
}





