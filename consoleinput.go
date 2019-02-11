package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(str, 64)

}
