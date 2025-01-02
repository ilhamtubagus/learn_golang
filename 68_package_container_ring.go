package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Implementation of circular linked list in golang

func main() {
	var data *ring.Ring = ring.New(3)
	data.Value = "Ilham"
	data.Next().Value = "Tubagus"
	data.Next().Next().Value = "Arfian"

	data.Do(func(v interface{}) {
		fmt.Println(v)
	})

	// or we can add using for loop

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(v interface{}) {
		fmt.Println(v)
	})
}
