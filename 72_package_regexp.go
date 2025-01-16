package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eco"))
	fmt.Println(regex.FindString("eko khanedi"))

	fmt.Println(regex.FindAllString("eko khanedi", -1))
	fmt.Println(regex.FindAllString("eko khanedi eco", 2))
}
