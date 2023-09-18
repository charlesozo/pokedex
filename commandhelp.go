package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, i string) error {
	if len(i) > 0{
		return errors.New("no such commands")
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}