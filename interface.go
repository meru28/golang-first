package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelo(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var joklik Person
	joklik.Name = "Bongklik"

	cat := Animal{
		Name: "Bela",
	}

	sayHelo(cat)

	sayHelo(joklik)
}
