package main

import "fmt"

func getGoodBye(nama string) string {
	return "Good bye " + nama
}

func main() {

	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("uplung"))
}
