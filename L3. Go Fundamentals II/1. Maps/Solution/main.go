package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[uint16]string{

		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		7: "English",
		8: "Cantonese",
	}

	for id, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(id, course)
		}
	}

	courses[4] = "Algorithms"
	courses[9] = "Spanish"

	delete(courses, 1)

	fmt.Println("Updated courses:", courses)
}
