package main

import (
	"fmt"
)

func findMinMax(numbers *[6]int) (int, int) {
	min := (*numbers)[0] // Ini adalah dereferencing
	max := (*numbers)[0] // Ini juga dereferencing

	for _, num := range *numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var numbers [6]int

	fmt.Println("Enter 6 numbers:")

	for i := 0; i < 6; i++ {
		fmt.Printf("Input %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	min, max := findMinMax(&numbers)

	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
}
