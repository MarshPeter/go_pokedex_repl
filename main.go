package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

	return commands
}

func commandExit() error {
	fmt.Println("Thank you for using the pokedex!")
	fmt.Println("Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------------")
	fmt.Printf("%-5vUsage Guidelines\n", "")
	fmt.Println()

	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
		
	}
	fmt.Println()
	fmt.Println("-----------------------------")

	return nil
}

func main() {
	var userInput string
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		fmt.Scanf("%s", &userInput)
		fmt.Println(userInput)
		fmt.Println()

		val, ok := commands[userInput]

		if !ok {
			fmt.Println("That command was not recongized.")
			fmt.Println("Please try again, or type 'help' to see a list of commands")
			continue
		}

		val.callback()
	}
}