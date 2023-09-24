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

var commands map[string]cliCommand // This is your package-level declaration

func init() {
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
	}
}

func commandHelp() error {
	fmt.Println("**********")
	fmt.Println("You asked for help!")
	fmt.Println("")
	fmt.Println("Here are a list of commands:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Print(cmd.name, ": ", cmd.description)
		fmt.Println("")
	}
	fmt.Println("**********")

	return nil
}

func commandExit() error {
	fmt.Println("**********")
	fmt.Println("Goodbye! You've exiting the Pokedex.")
	fmt.Println("**********")
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
