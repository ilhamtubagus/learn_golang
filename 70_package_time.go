package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current time
	var now time.Time = time.Now()
	fmt.Println("Now", now)

	// Get current time in local
	fmt.Println("Local time", now.Local())

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println("UTC time", utc)
	fmt.Println("Local time", utc.Local())

	// formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(time.DateTime, value)
	if err == nil {
		fmt.Println("Parsed time", valueTime)
	} else {
		fmt.Println("Error parsing time", err)
	}

	var duration time.Duration = time.Second * 10 // 10 seconds
	var duration2 time.Duration = time.Minute * 2 // 2 minutes
	var duration3 time.Duration = duration2 - duration
	fmt.Printf("Duration 1 %d\n", duration)
	fmt.Printf("Duration 2 %d\n", duration2)
	fmt.Printf("Duration 3 %d\n", duration3)
}
