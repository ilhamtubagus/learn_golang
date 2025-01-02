package main

import (
	"fmt"
	"learn_golang/database"
	_ "learn_golang/internal" // use underscore "-" to prevent error when importing unused package
)

func main() {
	fmt.Println(database.GetDatabase()) // Output: MySQL
}
