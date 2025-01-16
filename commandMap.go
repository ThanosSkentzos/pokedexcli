package main

import (
	"fmt"
)

func commandMap(config *Config, parameters []string) error {
	mapJSON, err := getMap(config.Next, &config.cache)
	if err != nil {
		fmt.Println("Error while getting map:", err)
		return err
	}
	config.Previous = config.Next
	config.Next = mapJSON.Next
	if config.Next == "" {
		config.Next = config.BaseAreaURL
	}
	printNames(getAreaNames(mapJSON))
	return nil
}

func commandMapb(config *Config, parameters []string) error {
	mapJSON, err := getMap(config.Previous, &config.cache)
	if err != nil {
		fmt.Println("Error while getting map:", err)
		return err
	}
	printNames(getAreaNames(mapJSON))
	return nil
}

func getAreaNames(mapJSON MapJSON) []string {
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
