package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%d", 1000000000)
	time.Sleep(1 * time.Second)
	for i := 10; i >= 0; i-- {
		fmt.Printf("\033[2K\r")
		time.Sleep(1 * time.Second)
		fmt.Printf("%d", i)
	}
	fmt.Println()
}
