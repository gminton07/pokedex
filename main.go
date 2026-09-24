package main

import (
	"fmt"
	"bufio"
	"os"
	"time"

	"github.com/gminton07/pokedex/internal/pokecache"
)

func main() {
	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	// REPL loop
	// initialize config
	localConfig := Config{
		Commands: commandRegistry(),
		NextURL:  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		PrevURL:  "null",
		Cache:    pokecache.NewCache(30 * time.Second),
	}
	fmt.Println("Welcome to the Pokedex!")

	for {
		// read
		fmt.Print("Pokedex > ")

		var args []string

		for scanner.Scan() {
			continue
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}

		str := scanner.Text()
		command := cleanInput(str)

		// eval && print
		firstWord := command[0]
		if command, ok := localConfig.Commands[firstWord]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&localConfig)
			if err != nil {
				fmt.Errorf("error: Command %s, %w", firstWord, err)
			}
		}
		
		// Testing code
		//fmt.Printf("Config: %s\n", localConfig)
	}
}
