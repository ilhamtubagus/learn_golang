package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// address1 duplicated to address2, so changes made to address2 will not affect address1
	address2 := address1

	address2.City = "Jakarta"
	// address1 and address2 are independent structs
	fmt.Println(address1) // Output: Subang
	fmt.Println(address2) // Output: Jakarta

	address3 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address4 are pointers to the same struct address3
	address4 := &address3
	// changes made to address4 will affect address3
	address4.City = "Depok"
	fmt.Println(address3.City) // Output: Depok
	fmt.Println(address4.City) // Output: Depok
}
