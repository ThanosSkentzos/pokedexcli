package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		text := s.Text()
		userInput := strings.Fields(strings.ToLower(text))
		name := userInput[0]
		command, ok := getCommands()[name]
		if !ok {
			fmt.Printf("Unknown command.")
			continue
		}
		command.callback()
	}
}