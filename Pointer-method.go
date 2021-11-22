package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}
