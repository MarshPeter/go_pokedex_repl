package main

import (
	"fmt"
)

func main() {
	var userInput string
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&userInput)
		fmt.Println()

		val, ok := commands[userInput]

		if !ok {
			fmt.Println("That command was not recongized.")
			fmt.Println("Please try again, or type 'help' to see a list of commands")
			continue
		}

		val.callback()
		fmt.Println()
	}
}