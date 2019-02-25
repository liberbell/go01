package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Println("Connection open: %v\n", isConnected)
	doSomething()
	fmt.Println("Connection open: %v\n", isConnected)
}

func doSomething() {
	connect()
	fmt.Println("Deffering disconnect!")
	defer disconnect()
	fmt.Println("Connection open: %v\n", isConnected)
	fmt.Println("Doing something!")
}
