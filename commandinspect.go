package main

import (
	"fmt"
)

func commandInspect(c *config, poke_name string) error {
	data, ok := pokemonTrack[poke_name]
	if ok {
		fmt.Printf("name: %s\n", data.name)
		fmt.Printf("weight: %d\n", data.weight)
		fmt.Printf("height: %d\n", data.height)
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil
}
