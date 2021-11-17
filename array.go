package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Meru"
	names[1] = "Rendy"
	names[2] = "Wirawan"

	fmt.Println(names[0], names[1], names[2])

	var values = [3]int{
		80,
		90,
		70,
	}

	fmt.Println(values)
}
