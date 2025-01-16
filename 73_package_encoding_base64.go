package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encoded = base64.StdEncoding.EncodeToString([]byte("Ilham Tubagus"))

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(decoded))
}
