package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Meru",
		"address": "Rawabuntu",
	}

	person["job"] = "Programmer"

	fmt.Println(len(person))
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Meru"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
