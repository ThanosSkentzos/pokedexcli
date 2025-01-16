package main

import "fmt"

func commandInspect(config *Config, parameters []string) error {
	pokemonName := parameters[0]
	pokemon, exists := config.pokedex[pokemonName]
	if !exists {
		fmt.Printf("%s not caught yet...\n",pokemonName)
		return nil
	}
	printPokemon(pokemon)
	return nil
}

func printPokemon(pokemon PokemonJSON){

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _,s := range pokemon.Stats{
		fmt.Printf("  -%s: %d\n",s.Stat.Name,s.BaseStat)
	}
	fmt.Println("Types:")
	for _,t := range pokemon.Types{
		fmt.Printf("  -%s\n",t.Type.Name)
	}
}