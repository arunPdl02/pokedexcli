package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		comandName := words[0]

		fmt.Printf("Your command was: %s\n", comandName)
	}
}

func cleanInput(text string) []string {
	result := []string{}
	text = strings.ToLower(text)

	start := -1

	for i, char := range text {
		if char != ' ' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				result = append(result, text[start:i])
				start = -1 // Reset state
			}
		}
	}

	if start != -1 {
		result = append(result, text[start:])
	}

	return result
}
