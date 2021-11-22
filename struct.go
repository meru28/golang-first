package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hallo", name, "Namaku adalah", customer.Name)
}

func main() {
	var obj Customer
	obj.Name = "Meru"
	obj.Address = "Rawabuntu"
	obj.Age = 31

	uplung := Customer{
		Name:    "Rendy",
		Address: "Rawabuntu",
		Age:     32,
	}

	fmt.Println("ini obj dari struct", obj)
	fmt.Println("ini obj dari struct", uplung)

	uplung.sayHi("Jokik")

}
