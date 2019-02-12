package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicity typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)
	str2 := "An explicity typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Println("Equal? ", (lvalue == uvalue))
	fmt.Println("Equal Non-Case-Sensitive? ", strings.EqualFold(lvalue, uvalue))

	fmt.Println("contain exp?", strings.Contains(str1, exp))
}
