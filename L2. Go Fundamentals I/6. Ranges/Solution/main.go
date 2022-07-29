package main

// Uncomment the line below after implementing your reduce() function
import "fmt"

func reduce(list []int) int {
	sum := 0

	for _, num := range list {
		sum += num
	}

	return sum
}

func main() {
	// Uncomment the line below after implementing your reduce() function
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
