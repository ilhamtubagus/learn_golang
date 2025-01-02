package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedByPointer() { // recommended using pointer for better performance
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{Name: "Eko"}

	eko.Married()
	fmt.Println(eko.Name) // Output: Eko

	eko.MarriedByPointer()
	fmt.Println(eko.Name) // Output: Mr. Eko
}
