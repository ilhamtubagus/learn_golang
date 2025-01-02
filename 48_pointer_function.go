package main

import "fmt"

type Address struct {
	Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndonesiaUsingPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// In default function parameters are passed by value
	// We can use pointer function to modify the original parameter value

	address := Address{"Singapore"}
	ChangeAddressToIndonesia(address)
	fmt.Println(address.Country) // Output: Singapore, not changing

	ChangeAddressToIndonesiaUsingPointer(&address)
	fmt.Println(address.Country) // Output: Indonesia, changing
}
