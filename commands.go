package main

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


func initializeCommands() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists available locations",
			callback:    commandMap,
		},
	}
}
