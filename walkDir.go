package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path: ", root)
}