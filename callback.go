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
	// Base URLs
	AreaBaseURL 	string
	PokemonBaseURL  string
	// map/mapb
	NextURL 	string
	PrevURL 	string
	// command cache
	Cache           pokecache.Cache
	// User Pokedex
	Pokedex 	map[string]pokeapi.Pokemon
	// Add more parameters as need arises
}


type cliCommand struct {
	name 		string
	description 	string
	callback 	func(*Config, []string) error
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
		"explore": {
			name:        "explore",
			description: "List pokemon in map area. Arg: area_name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch pokemon. Arg: pokemon_name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show info about caught pokemon. Arg: pokemon_name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List pokemon currently in Pokedex",
			callback:    commandPokedex,
		},
		// Add new commands
	}
	return cliCommandMap
}

func commandExit(c *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, args []string) error {
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

	return nil
}

func commandMap(C *Config, args []string) error {
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

func commandMapB(C *Config, args []string) error {
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

func commandExplore(C *Config, args []string) error {
	// Show pokemon which exist in a map area
	
	if len(args) < 1 {
		return errors.New(`error: "explore" command requires "area_name" argument`)
	}

	url := C.AreaBaseURL + args[0]
	err := pokeapi.GetAreaPokemon(url, &C.Cache)
	if err != nil {
		return err
	}
	
	return nil
}

func commandCatch(C *Config, args []string) error {
	// Catch chosen pokemon

	if len(args) < 1 {
		return errors.New(`error: "explore" command requires "pokemon_name" argument`)
	}

	pokemonName := args[0]

	// Check for pokemon in C.Pokedex
	if _, ok := C.Pokedex[pokemonName]; ok {
		fmt.Printf("Already have %s in pokedex\n", pokemonName)
		return nil
	}

	// Otherwise
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := C.PokemonBaseURL + pokemonName
	pokemon, err := pokeapi.GetPokemon(url)
	if err != nil {
		return err
	}
	if pokemon.Name == "" {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	
	fmt.Printf("%s was caught!\n", pokemonName)
	C.Pokedex[pokemonName] = pokemon
	
	return nil

}

func commandInspect(C *Config, args []string) error {
	// Show info about pokemon in the pokedex

	if len(args) < 1 {
		return errors.New(`error: "inspect" command requires "pokemon_name" argument`)
	}

	// Check for pokemon in C.Pokedex
	pokemonName := args[0]
	v, ok := C.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught that pokemon\n")
		return nil
	}

	// Print information
	fmt.Printf("Name: %s\n", v.Name)
	fmt.Printf("ID: %d\n", v.ID)
	fmt.Printf("Height: %d\n", v.Height)
	fmt.Printf("Weight: %d\n", v.Weight)
	fmt.Println("Stats:")
	for _, s := range v.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range v.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(C *Config, args []string) error {
	// List caught pokemon

	fmt.Println("You have these pokemon:")

	for k, _ := range C.Pokedex {
		fmt.Printf("  - %s\n", k)
	} 

	return nil
}
