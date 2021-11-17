package main

import "fmt"

//import "fmt"

func sayHello(ulang int) (uclung, kumplung int) {
	uclung = ulang
	kumplung = 3
	return
}

func main() {
	//for i := 0; i < 10; i++ {
	//	sayHello(i)
	//}
	//sayHello(10)
	a, b := sayHello(34)
	fmt.Println(a)
	fmt.Println(b)
}
