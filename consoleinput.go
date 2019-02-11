package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")

	str, _ := reader.ReadString('\n')
	fmt.Println(str)

}
