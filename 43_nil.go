package main

import "fmt"

// Only used in a couple of data types

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	m := NewMap("Ilham")
	fmt.Println(m["name"])
	m = NewMap("")

	if m == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println("Map is not nil")
	}
}
