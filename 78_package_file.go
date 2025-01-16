package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	_, err = file.WriteString(message)

	if err != nil {
		return err
	}

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}

	return message, nil
}

func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	_, err = file.WriteString(message)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	//err := createNewFile("sample.log", "Sample log")
	//if err != nil {
	//	return
	//}

	result, _ := readFile("sample.log")
	fmt.Println(result)

	_ = addToFile("sample.log", "New log entry\n")
}
