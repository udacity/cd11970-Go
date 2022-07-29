package main

import (
	"fmt"
	"time"
)

func iterateAndPrint(slice []string) {
	for _, item := range slice {
		fmt.Println(item)
	}
}

func main() {
	startTime := time.Now()

	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}

	// Run this program with and without the "go" keyword to see the difference in elapsed time for execution!
	go iterateAndPrint(colorNames)
	go iterateAndPrint(colorCodes)

	duration := time.Since(startTime)

	fmt.Println("Duration of execution: " + duration.String())

	time.Sleep(time.Second)
}
