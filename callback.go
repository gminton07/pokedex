package main

import (
	"fmt"
	"os"
	"errors"

	"github.com/gminton07/pokedex/internal/pokeapi"
)

type config struct{
	commands 	map[string]cliCommand
	// map/mapb
	nextURL 	string
	prevURL 	string
	// Add more parameters as need arises
}


type cliCommand struct {
	name 		string
	description 	string
	callback 	func(*config) error
}

func commandRegistry() map[string]cliCommand {
	cliCommandMap := map[string]cliCommand{
		"help": {
			name:       "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		// map commands
		"map": {
			name:        "map",
			description: "List next 20 map areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 map areas",
			callback:    commandMapB,
		},
		// Add new commands
	}
	return cliCommandMap
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	cliCommandMap := c.commands
	if len(cliCommandMap) < 1 {
		err := errors.New("error: no commands in registry")
		return err
	}

	// Print intro
	fmt.Println("Usage:\n")

	// loop through map
	for _, v := range cliCommandMap {
		fmt.Printf("%s:\t%s\n", v.name, v.description)
	}
	fmt.Println()

	return nil
}

func commandMap(c *config) error {
	// Show next 20 locations
	data, err := pokeapi.MapGet(c.nextURL)
	if err != nil {
		return err
	}

	// update config
	c.nextURL = data.Next
	c.prevURL = data.Previous

	return nil

}

func commandMapB(c *config) error {
	// Show previous 20 locations
	data, err := pokeapi.MapGet(c.prevURL)
	if err != nil {
		return err
	}
	//if len(data.Areas) == 0 {
	//	//fmt.Println("Empty struct returned")
	//	return nil
	//}

	// update config
	fmt.Printf("Next: %s\n", data.Next)
	fmt.Printf("Previous: %s\n", data.Previous)

	c.nextURL = data.Next
	if len(data.Previous) > 0 {
		c.prevURL = data.Previous
	}

	return nil
}
