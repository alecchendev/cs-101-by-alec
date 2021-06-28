package main

import (
	"fmt"
)

func main() {
	// playHangman()
	classifyEpicness()
}

func classifyEpicness() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)

	if input == "Alec" {
		fmt.Print("You're ")
		for i := 0; i < 3; i++ {
			fmt.Print("really ")
		}
		fmt.Print("epic.")
	} else if input == "Joe" {
		fmt.Print("You're ")
		for i := 0; i < 2; i++ {
			fmt.Print("really ")
		}
		fmt.Print("epic.")
	} else {
		fmt.Print("You're epic.")
	}
}
