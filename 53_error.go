package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi harus bernilai lebih besar dari 0")
	}

	return nilai / pembagi, nil
}

func main() {
	value, err := Pembagian(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
