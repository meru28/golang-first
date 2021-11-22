package main

import "fmt"

func endApp() {
	fmt.Println("Defer Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan recover message", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Meru")
}
