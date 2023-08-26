package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int)

	for _, item := range slice {
		countMap[item]++
	}

	return countMap
}

func main() {
	result := Mapping([]string{"asd", "qweqw", "asd", "adi"})
	fmt.Println(result) // Output: map[adi:1 asd:2 qweqw:1]
}
