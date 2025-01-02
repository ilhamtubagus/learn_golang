package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// Implement sort interface

// Len implements Len method for sorting
func (s UserSlice) Len() int { // No need asterisk (*) because Slice default is a pointer
	return len(s)
}

// Less implements Less method for sorting
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Swap implements Swap method for sorting
func (s UserSlice) Swap(i, j int) {
	//Old way to swap
	//temp := s[i]
	//s[i] = s[j]
	//s[j] = temp

	//New way to swap using Go's in-built swap function
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		User{Name: "Ilham", Age: 30},
		User{Name: "Tubagus", Age: 25},
		User{Name: "Arfian", Age: 28},
	}

	// Sorting users by age
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
