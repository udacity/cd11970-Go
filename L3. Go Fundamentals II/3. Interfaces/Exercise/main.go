package main

import "fmt"

type Rectangle struct {
	length, width float64
}

type Square struct {
	side float64
}

func (r Rectangle) rectanglePerimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (s Square) squarePerimeter() float64 {
	return s.side * 4
}

func main() {
	rectangle := Rectangle{length: 2, width: 4}
	square := Square{side: 2}

	fmt.Println("Rectangle perimeter: ", rectangle.rectanglePerimeter())
	fmt.Println("Square perimeter: ", square.squarePerimeter())
}
