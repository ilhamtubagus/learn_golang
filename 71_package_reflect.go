package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Sample struct {
	Name string `required:"true" max:"50"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type of value:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func main() {
	sample := Sample{Name: "Eko"}
	readField(sample)

	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println("Required field:", required)

	maxLength, _ := strconv.ParseInt(structField.Tag.Get("max"), 10, 10)
	fmt.Println("Max length:", maxLength)
}
