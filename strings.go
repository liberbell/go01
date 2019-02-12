package main

import (
	"fmt"
)

func main() {
	str1 := "An implicity typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)
	str2 := "An explicity typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)
}
