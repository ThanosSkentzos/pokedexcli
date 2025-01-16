package main

import "fmt"

func commandExplore(config *Config, parameters []string) error {
	baseUrl := config.BaseAreaURL
	area := parameters[0]
	url := baseUrl + area
	areaJSON, err := getArea(url, &config.cache)
	if err != nil {
		fmt.Println("Error while getting area information:", err)
		return err
	}
	printNames(getPokemonNames(areaJSON))
	return nil
}

func getPokemonNames(areaJSON AreaJSON) []string {
	var names []string
	for _, result := range areaJSON.PokemonEncounters {
		names = append(names, result.Pokemon.Name)
	}
	return names
}
