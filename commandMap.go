package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	mapJSON, err := getMAP(config.Next, &config.cache)
	if err != nil {
		fmt.Println("Error while getting map:", err)
		return err
	}
	config.Previous = config.Next
	config.Next = mapJSON.Next
	if config.Next == "" {
		config.Next = config.Initial
	}
	printNames(getNames(mapJSON))
	return nil
}

func commandMapb(config *Config) error {
	mapJSON, err := getMAP(config.Previous, &config.cache)
	if err != nil {
		fmt.Println("Error while getting map:", err)
		return err
	}
	printNames(getNames(mapJSON))
	return nil
}

func getNames(mapJSON MapJSON) []string {
	var names []string
	for _, result := range mapJSON.Results {
		names = append(names, result.Name)
	}
	return names
}

func printNames(names []string) {
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}
