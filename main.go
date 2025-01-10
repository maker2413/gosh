package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func rewrite(s string) {
	fmt.Printf("\033[2K\r")
	for i := 0; i <= len(s); i++ {
		fmt.Print(" ")
	}
	fmt.Printf("\033[2K\r")
	fmt.Printf(":> %s", s)
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
}

func getCommand() string {
	var line string
	fmt.Print(":> ")

	for {
		k := getKey()
		switch k[0] {
		case 2:
			// C-b
			fmt.Print("◀")
			line += "◀"
		case 6:
			// C-f
			fmt.Print("▶")
			line += "▶"
		case 13:
			// Enter
			return line
		case 16:
			// C-p
			fmt.Print("▲")
			line += "▲"
		case 14:
			// C-n
			fmt.Print("▼")
			line += "▼"
		case 27:
			switch k[1] {
			case 79:
				// Arrow key check
				switch k[2] {
				case 65:
					fmt.Print("▲")
					line += "▲"
				case 66:
					fmt.Print("▼")
					line += "▼"
				case 67:
					fmt.Print("▶")
					line += "▶"
				case 68:
					fmt.Print("◀")
					line += "◀"
				}
			}
		case 127:
			// Backspace
			if len(line) > 0 {
				line = line[0 : len(line)-1]
			}

			rewrite(line)
		default:
			fmt.Print(string(k[0]))
			line += string(k[0])
		}
	}
}

func readUserInput() {
	for {
		command := getCommand()
		if command == "exit" {
			return
		} else {
			fmt.Printf("\n%s\n", command)
		}
	}
}

func main() {
	readUserInput()

	// fmt.Println(getKey())

	// rewrite()
}
