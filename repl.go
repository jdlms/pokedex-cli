package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := strings.ToLower(scanner.Text())
		if cmd, exists := commands[line]; exists {
			err := cmd.callback(myConfig)
			if err != nil {
				fmt.Println("There was an error.", err)
			}

		} else {
			fmt.Println("Unknown command. Try again, or ask for help.")
		}
	}

}
