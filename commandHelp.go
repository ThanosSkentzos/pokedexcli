package main

import "fmt"

func commandHelp() error {
	println("Welcome to the Pokedex!")
	println("Usage:")

	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
