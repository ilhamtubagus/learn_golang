package main

type Address struct {
	Country string
}

func main() {
	address1 := new(Address)
	address2 := address1
	address1.Country = "Indonesia"
	println(address1.Country) // Output: Indonesia
	println(address2.Country) // Output: Indonesia
}
