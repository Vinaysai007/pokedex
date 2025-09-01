package main

import "strings"

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowerText := strings.ToLower(trimmed)
	words := strings.Fields(lowerText)
	return words
}
