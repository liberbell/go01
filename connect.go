package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Println("Connection open: %v\n", isConnected)
	doSomething()
	fmt.Println("Connection open: %v\n", isConnected)
}
