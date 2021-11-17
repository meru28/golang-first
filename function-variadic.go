package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	total := sumAll(10, 12, 23, 34)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	total = sumAll(slice...)
	fmt.Println(total)
}
