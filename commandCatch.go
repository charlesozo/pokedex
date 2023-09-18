package main

import (
	"fmt"
	"math/rand"
)

type poke struct {
	name   string
	height int
	weight int
}

var pokemonTrack map[string]poke = make(map[string]poke)

func commandCatch(c *config, poke_name string) error {
	resp, err := c.pokeapiClient.PokeData(poke_name)
	if err != nil {
		return err
	}
	_, exists := pokemonTrack[poke_name]
	if exists {
		fmt.Printf("\n%s has already been caught\n", poke_name)
	} else {
		fmt.Printf("Throwing a Pokeball at %s...", poke_name)
		baseExperience := resp.BaseExperience
		weightCaught := float32(baseExperience) / 200.0
		randomValue := rand.Float32()
		if weightCaught > randomValue {
			fmt.Printf("\n%s was caught\n", poke_name)
			pokemonTrack[poke_name] = poke{
				name:   poke_name,
				height: resp.Height,
				weight: resp.Weight,
			}
		} else {
			fmt.Printf("\n%s escaped\n", poke_name)
		}
	}
	return nil
}
