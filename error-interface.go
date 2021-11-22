package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai float32, pembagi float32) (float32, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// var contohError error = errors.New("Ups Error")
	hasil, err := pembagi(100, 8)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
	// fmt.Println(contohError)
}
