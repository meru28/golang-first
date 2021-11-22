package main

import "fmt"

func logging() {
	fmt.Println("defer memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("Hasil", result)
}

func main() {
	runApplication(10)
}
