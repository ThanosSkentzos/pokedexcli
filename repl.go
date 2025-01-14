package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ThanosSkentzos/pokedexcli/internal/pokecache"
)

type Config struct {
	Previous string
	Next     string
	Initial  string
	cache    pokecache.Cache
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.FieldsFunc(text, Separators)

	return result
}

func Separators(r rune) bool {
	return strings.ContainsRune(" ,.", r)
}

func startREPL() {
	s := bufio.NewScanner(os.Stdin)
	config := Config{
		"https://pokeapi.co/api/v2/location-area/",
		"https://pokeapi.co/api/v2/location-area/",
		"https://pokeapi.co/api/v2/location-area/",
		pokecache.NewCache(1000),
	}

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		text := s.Text()
		userInput := strings.Fields(strings.ToLower(text))
		if len(userInput) < 1 {
			continue
		}
		name := userInput[0]
		command, ok := getCommands()[name]
		if !ok {
			fmt.Printf("Unknown command.\n")
			continue
		}
		command.callback(&config)
	}
}
