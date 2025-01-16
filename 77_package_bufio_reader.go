package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	// buffered io, to create data IO Reader and Writer
	input := strings.NewReader("Hello, World!\nMy name is Ilham")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		// End of file then break the loop
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	input = strings.NewReader("1,2,3,4,5")

	reader = bufio.NewReader(input)

	for {
		char, err := reader.ReadString(',')

		if err == io.EOF {
			break
		}

		fmt.Println(strings.TrimSuffix(string(char), ","))
	}
}
