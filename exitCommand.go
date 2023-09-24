package main

import (
	"fmt"
	"os"
)

func commandExit(*config) error {
	fmt.Println("**********")
	fmt.Println("Goodbye! You've exiting the Pokedex.")
	fmt.Println("**********")
	os.Exit(0)
	return nil
}
