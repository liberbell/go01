package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path: ", root)

	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
