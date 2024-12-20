package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

func rewrite() {
	fmt.Printf("%d", 1000000000)
	time.Sleep(1 * time.Second)
	for i := 10; i >= 0; i-- {
		fmt.Printf("\033[2K\r")
		time.Sleep(1 * time.Second)
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func getKey() []byte {
	// Get the input file descriptor
	fd := int(os.Stdin.Fd())

	// Save the original terminal state to restore later
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	// Create a new terminal instance
	//term := term.NewTerminal(os.Stdin, "")

	b := make([]byte, 3)
	_, err = os.Stdin.Read(b)
	if err != nil {
		panic(err)
	}

	return b

	// for {
	// 	// Read a single byte
	// 	b, err := term.ReadByte()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	switch b {
	// 	case 27: // ESC
	// 		b2, _ := term.ReadByte()
	// 		if b2 == 91 { // '['
	// 			b3, _ := term.ReadByte()
	// 			switch b3 {
	// 			case 65:
	// 				fmt.Println("Up arrow")
	// 			case 66:
	// 				fmt.Println("Down arrow")
	// 			case 67:
	// 				fmt.Println("Right arrow")
	// 			case 68:
	// 				fmt.Println("Left arrow")
	// 			}
	// 		}
	// 	default:
	// 		fmt.Printf("Key pressed: %v\n", b)
	// 	}
	// }
}

func readUserInput() {
	var line strings.Builder
	fmt.Print(":> ")

	for {
		k := getKey()
		if k[0] == byte(13) {
			fmt.Println(line)
			break
		} else {
			fmt.Print(string(k))
			line.WriteString(string(k))
		}
	}
}

func main() {
	readUserInput()

	// rewrite()

}
