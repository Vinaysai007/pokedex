package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		first := words[0]
		fmt.Printf("Your command was: %s\n", first)
	}
}
