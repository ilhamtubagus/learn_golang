package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"ID cannot be empty"}
	}

	if id != "fian" {
		return &NotFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", "Hello, World!")
	//if err != nil {
	//	switch err.(type) {
	//	case *ValidationError:
	//		println(err.Error())
	//	case *NotFoundError:
	//		println(err.Error())
	//	default:
	//		println("Unknown error")
	//	}
	//}
	if err != nil {
		if validationError, ok := err.(*ValidationError); ok {
			fmt.Println("validation error: ", validationError.Error())
		} else if notFoundError, ok := err.(*NotFoundError); ok {
			fmt.Println("not found error: ", notFoundError.Error())
		} else {
			fmt.Println("unknown error: ", err.Error())
		}
	}
}
