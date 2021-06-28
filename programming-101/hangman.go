package main

import (
	"fmt"
	"syscall"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
)

func playHangman() {
	fmt.Println("Welcome to hangman!")
	fmt.Print("Provide a word to guess: ")

	secretBytes, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	secret := string(secretBytes)

	guessedWord := strings.Repeat("_", len(secret))

	guessedChars := []string{}

	for guessedWord != secret {
		fmt.Println("To guess:", guessedWord)
		fmt.Println("Guessed:", guessedChars)
		
		fmt.Print("Take a guess: ");
		var guess string
		fmt.Scanln(&guess)

		if len(guess) == len(secret) {
			if guess == secret {
				break
			}
		} else if len(guess) != 1 {
			fmt.Println("Please enter either a single character or a full word.")
			continue
		}

		for pos, char := range secret {
			if string(char) == guess {
				newGuessed := []rune(guessedWord)
				newGuessed[pos] = char
				guessedWord = string(newGuessed)
			}
		}

		guessedChars = append(guessedChars, guess)
	}

	fmt.Println("Congrats! You got the word! ->", secret)
	
}