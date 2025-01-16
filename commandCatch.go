package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, parameters []string) error {
	baseUrl := config.BasePokemonURL
	pokemonName := parameters[0]
	url := baseUrl + "/" + pokemonName
	pokemonJSON, err := getPokemon(url, &config.cache)
	if err != nil {
		fmt.Println("Error while getting pokemon:", err)
		return err
	}
	caught := catchPokemon(pokemonJSON)
	// ADD TO POKEDEX map[string]PokemonJSON
	if caught {
		config.pokedex[pokemonName] = pokemonJSON
	}
	return nil
}

func catchPokemon(pokemon PokemonJSON) bool {
	fmt.Printf("Throwing a Pokeball at %s...", pokemon.Name)
	baseExp := pokemon.BaseExperience
	resistance := rand.Intn(baseExp)
	if resistance < 42 { //CAUGHT
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return true
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return false
}
