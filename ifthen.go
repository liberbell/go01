package main

import "fmt"

func main() {
	var x float64 = 42
	var result string

	if x < 0 {
		result = "Less than zero"
	} else {
		result = "Greater than or equal to zero"
	}
	fmt.Println("Result:", result)

}
