package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("haloo")
	error := errors.New("Error bleng")
	fmt.Println(error.Error())

}
