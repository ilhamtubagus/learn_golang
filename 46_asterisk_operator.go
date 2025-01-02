package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Bandung"
	//fmt.Println(address1) // changed to Bandung
	//fmt.Println(address2) // changed to Bandung but the type is pointer

	//address2 = &Address{"Jakarta", "Jawa Barat", "Indonesia"}
	//fmt.Println(address1) // unchanged to Bandung
	//fmt.Println(address2) // changed to Jakarta but the type is pointer

	*address2 = Address{"Jakarta", "Jawa Barat", "Indonesia"}
	fmt.Println(address1) // changed to Jakarta
	fmt.Println(address2) // changed to Jakarta but the type is pointer
}
