package main

import "fmt"

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "purple")
	fmt.Println(colors)
}
