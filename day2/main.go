package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func isMultipleOf7(number int) bool {
	return number%7 == 0
}

func trapezoidArea(base1, base2, height float64) float64 {
	return (base1 + base2) * height / 2
}

func main() {
	var choice int
	fmt.Println("1. Menentukan Bilangan Prima")
	fmt.Println("2. Menentukan Bilangan Kelipatan 7")
	fmt.Println("3. Menghitung Luas Trapesium")
	fmt.Print("Pilih fungsi (1/2/3): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var number int
		fmt.Print("Masukkan angka: ")
		fmt.Scanln(&number)
		if isPrime(number) {
			fmt.Println("Bilangan prima")
		} else {
			fmt.Println("Bukan bilangan prima")
		}
	case 2:
		var number int
		fmt.Print("Masukkan angka: ")
		fmt.Scanln(&number)
		if isMultipleOf7(number) {
			fmt.Println("Bilangan merupakan kelipatan 7")
		} else {
			fmt.Println("Bilangan bukan kelipatan 7")
		}
	case 3:
		var base1, base2, height float64
		fmt.Print("Masukkan panjang alas bawah: ")
		fmt.Scanln(&base1)
		fmt.Print("Masukkan panjang alas atas: ")
		fmt.Scanln(&base2)
		fmt.Print("Masukkan tinggi trapesium: ")
		fmt.Scanln(&height)
		area := trapezoidArea(base1, base2, height)
		fmt.Printf("Luas trapesium: %.2f\n", area)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
