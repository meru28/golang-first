package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Banteng", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jatim", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Tangsel",
		Province: "Banten",
		Country:  "",
	}
	alamatPointer := &alamat
	changeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
