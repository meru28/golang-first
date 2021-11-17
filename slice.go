package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	//var slice =  months[4:7]

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
	//fmt.Println(slice)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Meru"
	newSlice[1] = "Rendy"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//
	//months[5] = "Diubah"
	//fmt.Println(slice)
	//
	//slice[0] = "Mei Lagi"
	//fmt.Println(months)
	//
	//days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	//daySlice1 := days[5:]
	//daySlice1[0] = "Sabtu Baru"
	//daySlice1[1] = "Minggu Baru"
	//fmt.Println(days)
	//
	//daySlice2 := append(daySlice1, "Libur Baru")
	//daySlice2[0] = "Ups"
	//fmt.Println(daySlice2)
	//fmt.Println(days)

}
