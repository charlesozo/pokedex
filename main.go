package main

import (
	"github.com/charlesozo/pokedex/internal/api"
	"time"
)


func main() {
	pokeClient := api.NewClient(5 * time.Second, time.Hour)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
	
}
