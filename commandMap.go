package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	mapJSON, err := getMAP(config.Next)
	if err != nil {
		fmt.Println("Error while getting map:", err)
		return err
	}
	config.Previous = config.Next
	config.Next = mapJSON.Next
	names := getNames(mapJSON)

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
	return nil
}

func getNames(mapJSON MapJSON) []string {
	var names []string
	for _, result := range mapJSON.Results {
		names = append(names, result.Name)
	}
	return names
}
