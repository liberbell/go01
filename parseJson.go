package main

import "fmt"

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
