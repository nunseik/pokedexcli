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
		fmt.Println("Your command was:", words[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}