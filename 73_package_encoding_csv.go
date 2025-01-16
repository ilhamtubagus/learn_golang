package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Read CSV data from a string
	csvString := "name,age\n" +
		"Eko,25\n" +
		"John,30\n" +
		"Ilham,26"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		// End of file then break the loop
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	// Write CSV data to a stdOut
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"name", "age"})
	_ = writer.Write([]string{"Eko", "25"})
	_ = writer.Write([]string{"John", "30"})
	_ = writer.Write([]string{"Ilham", "26"})
	writer.Flush()
}
