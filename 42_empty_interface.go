package main

import "fmt"

func Ups() any {
	return 1
}

func main() {
	var empty any = Ups()
	fmt.Println(empty) // Output: 1
}
