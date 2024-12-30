package main

import (
	"flag"
	"fmt"
)

func main() {
	// For parse command line arguments
	host := flag.String("host", "localhost", "hostname to connect to")
	username := flag.String("username", "user", "username for authentication")
	password := flag.String("password", "", "password for authentication")

	flag.Parse()

	fmt.Printf("Connecting to %s as %s [%s]...\n", *host, *username, *password)
}
