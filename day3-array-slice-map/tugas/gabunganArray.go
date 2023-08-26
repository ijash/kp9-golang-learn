package main

import "fmt"

func arrayMerge(arrA, arrB []string) []string {
	merged := append(arrA, arrB...)
	uniqueMap := make(map[string]bool)
	result := []string{}

	for _, name := range merged {
		if !uniqueMap[name] {
			uniqueMap[name] = true
			result = append(result, name)
		}
	}

	return result
}

func main() {
	arrayA := []string{"Aulia", "Budi", "Citra", "Dian", "Eka"}
	arrayB := []string{"Budi", "Faisal", "Citra", "Gita", "Dian"}

	mergedArray := arrayMerge(arrayA, arrayB)

	fmt.Println("Merged Array with Unique Names:", mergedArray)
}
