package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config, i string) error {
	if len(i) > 0{
		return errors.New("no such commands")
	}
	if len(pokemonTrack)==0{
		fmt.Println("you haven't caught any pokemon")
	}else{
		for k:= range pokemonTrack{
			fmt.Printf("- %s\n", k)
		}
	}
	return nil
}