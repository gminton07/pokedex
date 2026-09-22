package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	// REPL loop
	// initialize config
	localConfig := config{
		commands: commandRegistry(),
		nextURL:  "https://pokeapi.co/api/v2/location-area/",
		prevURL:  nil,
	}
	fmt.Println("Welcome to the Pokedex!")

	for {
		// read
		fmt.Print("Pokedex > ")

		done := scanner.Scan()
		if done == false {
			if err := scanner.Err; err != nil {
				fmt.Errorf("error: %w", err)
			}	
		}

		str := scanner.Text()
		command := cleanInput(str)

		// eval && print
		firstWord := command[0]
		if command, ok := localConfig.commands[firstWord]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&localConfig)
			if err != nil {
				fmt.Errorf("error: Command %s, %w", firstWord, err)
			}
		}
	}
}
