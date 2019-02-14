package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.February, 15, 8, m33, 0, 0, time.Jst)
	fmt.Printf("Go launched at %s\n", t)
}
