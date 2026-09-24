package main

import (
	"fmt"
	"os"
	"errors"

	"github.com/gminton07/pokedex/internal/pokeapi"
	"github.com/gminton07/pokedex/internal/pokecache"
)

type Config struct{
	Commands 	map[string]cliCommand
	// map/mapb
	NextURL 	string
	PrevURL 	string
	// command cache
	Cache           pokecache.Cache
	// Add more parameters as need arises
}


type cliCommand struct {
	name 		string
	description 	string
	callback 	func(*Config) error
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

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	cliCommandMap := c.Commands
	if len(cliCommandMap) < 1 {
		err := errors.New("error: no commands in registry")
		return err
	}

	// Print intro
	fmt.Println("Usage:")

	// loop through map
	for _, v := range cliCommandMap {
		fmt.Printf("%s:\t%s\n", v.name, v.description)
	}
	fmt.Println()

	return nil
}

func commandMap(C *Config) error {
	// Show next 20 locations
	data, err := pokeapi.MapGet(C.NextURL, &C.Cache)
	if err != nil {
		return err
	}

	// update config
	C.NextURL = data.Next
	C.PrevURL = data.Previous

	return nil

}

func commandMapB(C *Config) error {
	// Show previous 20 locations
	data, err := pokeapi.MapGet(C.PrevURL, &C.Cache)
	if err != nil {
		return err
	}
	if len(data.Results) == 0 {
		// Check if data is empty
		return nil
	}

	// update config
	C.NextURL = data.Next
	C.PrevURL = data.Previous

	return nil
}
