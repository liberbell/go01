package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Sum: ", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 100000; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)

}