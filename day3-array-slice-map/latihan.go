package main

import "fmt"

func main() {
	//array
	var arr [5]int
	var arr2 [5]int
	arr2 = [5]int{5, 3, 7, 9, 5}
	arr3 := [5]int{5, 3, 7, 9, 5}

	//slice
	sliceData2 := make(map[string]int) // langsung isi data
	fmt.Println(sliceData2)
	var sliceData []int
	sliceData = append(sliceData, 10)
	sliceData = append(sliceData, []int{5, 3, 8}...) // syaratnya harus ada titik 3
	sliceData = append([]int{2, 3, 3}, sliceData...) // kalau mau tambah di depan
	// cara delete manual
	sliceData = sliceData[:6] // intinya ngambil dari berapa: ke berapa, bisa dikosongkan.
	fmt.Println(arr, arr2, arr3)
	fmt.Println(sliceData)

	for i, v := range sliceData {
		fmt.Printf("index: %d nilai:%d\n", i, v)
	}

	//maps
	var tabunganSiswa map[string]int = map[string]int{}
	tabunganSiswa["heru"] = 10000
	tabunganSiswa["budi"] = 12500
	fmt.Println(tabunganSiswa["heru"])
	hasilMap, isFound := tabunganSiswa["gilang"]
	fmt.Printf("key ketemu?%t , isinya: %d\n", isFound, hasilMap)
	for key, value := range tabunganSiswa {
		fmt.Println(key, value) // jadi bukan index, ga ngambil sesuai urutan, yg penting kepanggil.
	}

}
