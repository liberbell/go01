package main

import "io/ioutil"

func main() {
	fileName := "hello.txt"

	content, err := ioutil.ReadFile(fileName)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
