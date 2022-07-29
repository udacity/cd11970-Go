package main

import "fmt"

func getRectangleArea(width, length int) string {
	product := width * length

	if product < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", product)
	} else {
		return fmt.Sprintf("The area is %d, which is greater or equal to 50", product)
	}
}

func main() {
	fmt.Println(getRectangleArea(5, 10))
}
