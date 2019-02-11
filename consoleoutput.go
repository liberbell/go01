package main

import (
	"fmt"
)

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	// stringLength, err := fmt.Println(str1, str2, str3)
	stringLength, _ := fmt.Println(str1, str2, str3)

	// if err == nil {
	fmt.Println("String length:", stringLength)
	// }

	fmt.Printf("value of number: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
}
