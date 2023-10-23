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
		"mapb": {
			name: "mapb",
			description: "Get the previous 20 locations",
			callback: commandMapB,
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

var currentLocation locations

func commandMap() error {
	if currentLocation.Next == "" {
		currentLocation = getNextMapPage("https://pokeapi.co/api/v2/location/")
	} else if (currentLocation.Next != "null") {
		currentLocation = getNextMapPage(currentLocation.Next) 
	} else {
		fmt.Println("There are no more pages.")
		return nil
	}

	for _, value := range currentLocation.Results {
		fmt.Println(value.Name)
	} 

	return nil
}

func commandMapB() error {
	if currentLocation.Previous == "" || currentLocation.Previous == "null" {
		fmt.Println("You are already on the first page.")
		return nil
	}

	currentLocation = getNextMapPage(currentLocation.Previous)

	for _, value := range currentLocation.Results {
		fmt.Println(value.Name)
	}

	return nil
}