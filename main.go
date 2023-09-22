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

// return map[string]cliCommand{
//     "help": {
//         name:        "help",
//         description: "Displays a help message",
//         callback:    commandHelp,
//     },
//     "exit": {
//         name:        "exit",
//         description: "Exit the Pokedex",
//         callback:    commandExit,
//     },
// }

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		if line == "quit" || line == "stop" {
			fmt.Println("bye!")
			break
		}

	}

}
