package main

import "fmt"

func commandHelp(*config) error {
	fmt.Println("**********")
	fmt.Println("You asked for help!")
	fmt.Println("")
	fmt.Println("Here are a list of commands:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Print(cmd.name, ": ", cmd.description)
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("**********")
	fmt.Println("")

	return nil
}
