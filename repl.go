package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct{
	Previous string
	Next string
}

func cleanInput(text string) []string{
	text = strings.ToLower(text)
	result := strings.FieldsFunc(text,Separators)

	return result
}

func Separators(r rune) bool{
	return strings.ContainsRune(" ,.",r)
}


func startREPL(){
	s := bufio.NewScanner(os.Stdin)
	config := Config{
		"",
		"https://pokeapi.co/api/v2/location-area/",
	}

	for {
		fmt.Print("Pokedex > ")
		
		// commandMap(&config)
		// commandMap(&config)
		// commandMap(&config)
		// commandMap(&config)
		s.Scan()
		text := s.Text()
		userInput := strings.Fields(strings.ToLower(text))
		name := userInput[0]
		command, ok := getCommands()[name]
		if !ok {
			fmt.Printf("Unknown command.")
			continue
		}
		command.callback(&config)
	}
}