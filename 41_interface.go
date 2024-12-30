package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

type Animal struct {
	name string
}

func (a Animal) GetName() string {
	return fmt.Sprintf("Animal %s", a.name)
}

func SayHello(value HasName) {
	fmt.Printf("Hello, %s!\n", value.GetName())
}

func main() {
	person := Person{name: "Ilham"}
	SayHello(&person)

	animal := Animal{name: "Lion"}
	SayHello(&animal)
}
