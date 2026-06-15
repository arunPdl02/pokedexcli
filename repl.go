package main

import "strings"

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
