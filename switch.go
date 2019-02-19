package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	case 1:
		result = "It`s sunday."
	case 7:
		result = "It`s saturday."
	default:
		result = "It`s weekday."

	}
	fmt.Println("Day", dow, ",", result)
}
