package main

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}


func getCommands() map[string]cliCommand{
	commands := map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map":{
				name:		 "map",
				description: "Displays the map",
				callback:	 commandMap,
			},
		}
	return commands
}