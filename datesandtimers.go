package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.February, 15, 8, 37, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("The month is ", t.Month())
	fmt.Println("The day is ", t.Day())
	fmt.Println("The weekday is ", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorro is %v, %v %v, %v\n",
		tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Friday, February 15, 2019"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))
}
