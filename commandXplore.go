package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	if len(location)==0 {
		return errors.New("you must provide a location Name")
	}
	resp, err := cfg.pokeapiClient.Xplore(location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n...", resp.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}