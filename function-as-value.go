package main

import (
	"fmt"
	"runtime"
)

func getGoodBye(nama string) string {
	return "Good bye " + nama
}

func main() {

	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("uplung"))
	fmt.Println(runtime.GOOS)
}
