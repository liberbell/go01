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
	fmt.Printf("value of number as float: %.2f\n", float64(aNumber))

	fmt.Printf("Data types       : %T,%T,%T,%T, and,%T\n",
		str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data types as var: %T,%T,%T,%T, and,%T\n",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
