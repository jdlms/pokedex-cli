package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
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
}

func commandHelp() error {
	fmt.Println("You asked for help!")
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye! You've exiting the Pokedex.")
	os.Exit(0)
	return nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		if cmd, exists := commands[line]; exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println("There was an erorr.", err)
			}

		} else {
			fmt.Println("Unknown command. Try again.")
		}
	}

}
