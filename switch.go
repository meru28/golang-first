package main

import "fmt"

func main() {
	//switch name := "Maghfiroh"; name {
	//case "Meru":
	//	fmt.Println("Halo Meru")
	//case "Hany":
	//	fmt.Println("Halo Hany")
	//case "Yuniar":
	//	fmt.Println("Halo Yuniar")
	//default:
	//	fmt.Println("Ngapain")
	//}
	length := len("Bucoks")
	switch  {
	case length > 5:
		fmt.Println("Nama Cukup")
	case length > 7:
		fmt.Println("Kepanjangan")
	default:
		fmt.Println("bener wes")
	}

}
