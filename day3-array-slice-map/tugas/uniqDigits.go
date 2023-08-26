package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Masukan beberapa nomer: ")
	fmt.Scanln(&input)

	result := findUniqueNumbers(input)
	fmt.Println("Output:", result)
}

func findUniqueNumbers(input string) []int {

	digitMap := make(map[int]int)
	uniqMap := make(map[int]int)
	for i := 0; i < len(input); i++ {
		// cara primitif untuk convert char -> int.
		// inget! di golang double quote, single quote ama backtick beda. https://golangbyexample.com/double-single-back-quotes-go/
		// ini karena urutan ascii nomer 0 yaitu 48
		digit := int(input[i]) - int('0') // bisa jg  int(input[i] - '0')

		digitMap[digit]++
		if digitMap[digit] > 1 {

			delete(uniqMap, digit)
		} else {
			uniqMap[digit]++
		}
		fmt.Println(digit)
	}
	uniqArr := make([]int, 0, len(uniqMap))

	for key, _ := range uniqMap {
		uniqArr = append(uniqArr, key)
	}
	return uniqArr
}
