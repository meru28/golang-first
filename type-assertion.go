package main

import "fmt"

func random() interface{} {
	return 5
}

func main() {
	result := random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("Type value", value, "is string")
	case int:
		fmt.Println("Type value", value, "is int")
	default:
		fmt.Println("unknown type")
	}
}
