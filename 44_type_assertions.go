package main

import (
	"fmt"
)

func random() any {
	return true
}

func main() {
	//result := random()
	//resultString := result.(string)
	//fmt.Println(resultString) // Output: OK

	// resultInt := result.(int) // This will panic
	// fmt.Println(resultInt)

	switch value := random().(type) {
	case string:
		fmt.Println("Result is a string", value)
	case int:
		fmt.Println("Result is an integer", value)
	default:
		fmt.Println("Result is of unknown type", value)
	}
}
