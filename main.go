package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	s := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		text := s.Text()
		userInput := strings.Fields(strings.ToLower(text))
		name := userInput[0]
		command, ok := commands[name]
		if !ok {
			fmt.Printf("Unknown command.")
			continue
		}
		command.callback()
	}
}
