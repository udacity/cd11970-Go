package main

import "fmt"

type Shape interface {
	perimeter() float64
}

type Rectangle struct {
	length, width float64
}

type Square struct {
	side float64
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	rectangle := Rectangle{length: 2, width: 4}
	square := Square{side: 2}

	fmt.Println("Rectangle perimeter: ", getPerimeter(rectangle))
	fmt.Println("Rectangle perimeter: ", getPerimeter(square))
}
