package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Eko", "John", "Ilham"}
	values := []int{25, 30, 26}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Index(names, "Ilham"))
}
