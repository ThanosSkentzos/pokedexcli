package main

import "fmt"

func commandPokedex(config *Config, parameters []string) error {
	for _,p:= range config.pokedex{
		fmt.Printf(" -%s\n",p.Name)
	}
	return nil
}