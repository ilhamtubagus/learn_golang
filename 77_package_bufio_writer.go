package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello, World!\n")
	err := writer.Flush()
	if err != nil {
		return
	} // flush the buffer to stdout
}
