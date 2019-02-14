package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.February, 15, 8, 37, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)
}
