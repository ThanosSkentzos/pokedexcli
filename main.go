package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		text := s.Text()
		words := strings.Fields(strings.ToLower(text))
		fmt.Printf("Your command was: %v\n",words[0])

	}
}

