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
		"map": {
			name: "map",
			description: "Get the next 20 locations",
			callback: commandMap,
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

func commandMap() error {
	getNextMapPage(20)
	return nil
}