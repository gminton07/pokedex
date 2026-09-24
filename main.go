package main

import (
	"fmt"
	"bufio"
	"os"
	"time"

	"github.com/gminton07/pokedex/internal/pokecache"
	"github.com/gminton07/pokedex/internal/pokeapi"
)

func main() {
	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	// REPL loop
	// initialize config
	localConfig := Config{
		Commands: commandRegistry(),
		AreaBaseURL:  "https://pokeapi.co/api/v2/location-area/",
		PokemonBaseURL: "https://pokeapi.co/api/v2/pokemon/",
		NextURL:  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		PrevURL:  "null",
		Cache:    pokecache.NewCache(30 * time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	fmt.Println("Welcome to the Pokedex!")

	for {
		// read
		fmt.Print("Pokedex > ")

		if done := scanner.Scan(); !done {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error:", err)
			}
		}

		str := scanner.Text()
		command := cleanInput(str)

		// eval && print
		cmd := command[0]
		args := command[1:]
		
		if command, ok := localConfig.Commands[cmd]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&localConfig, args)
			if err != nil {
				fmt.Println(fmt.Errorf("error: Command %s, %w", cmd, err))
			}
		}
		fmt.Println()
		
		// Testing code
		//fmt.Printf("Config: %s\n", localConfig)
	}
}
